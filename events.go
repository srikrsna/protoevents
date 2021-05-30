package protoevents

import (
	"context"
	"strings"
	"sync"

	evtpb "github.com/srikrsna/protoevents/events"
	"github.com/taskcluster/slugid-go/slugid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

//go:generate mockgen -package mocks -destination mocks/$GOFILE . Dispatcher,ErrorLogger

type Event struct {
	Id   string
	Type string

	Request  protoreflect.ProtoMessage
	Response protoreflect.ProtoMessage
}

type Dispatcher interface {
	Dispatch(context.Context, *Event) error
}

type ErrorLogger interface {
	Log(context.Context, *Event, error)
}

type Options struct {
	Dispatcher Dispatcher

	ErrorLogger ErrorLogger
}

func NewInterceptor(opt Options) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			return
		}

		evt := evtBuf.Get().(*Event)
		defer evtBuf.Put(evt)

		*evt = Event{
			Id:       "evt_" + slugid.Nice(),
			Type:     info.FullMethod,
			Request:  req.(protoreflect.ProtoMessage),
			Response: resp.(protoreflect.ProtoMessage),
		}

		mutex.RLock()
		v, ok := cache[info.FullMethod]
		mutex.RUnlock()
		if !ok {
			parts := strings.Split(info.FullMethod, "/")
			serviceName := protoreflect.FullName(parts[1])
			methodName := protoreflect.Name(parts[2])

			descriptor, derr := protoregistry.GlobalFiles.FindDescriptorByName(serviceName)
			if derr != nil {
				opt.ErrorLogger.Log(ctx, evt, derr)
				return
			}

			v = descriptor.(protoreflect.ServiceDescriptor).Methods().ByName(methodName).Options().ProtoReflect().Get(evtpb.E_Fire.TypeDescriptor()).Bool()

			mutex.Lock()
			cache[info.FullMethod] = v
			mutex.Unlock()
		}

		if !v {
			return
		}

		if err := opt.Dispatcher.Dispatch(ctx, evt); err != nil {
			opt.ErrorLogger.Log(ctx, evt, err)
		}

		return
	}
}

var cache = map[string]bool{}

var mutex sync.RWMutex

var evtBuf = sync.Pool{
	New: func() interface{} {
		return new(Event)
	},
}
