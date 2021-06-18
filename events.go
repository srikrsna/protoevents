package protoevents

import (
	"context"
	"strings"
	"sync"

	evtpb "github.com/srikrsna/protoevents/events"
	"github.com/taskcluster/slugid-go/slugid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
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

	Preprocess bool
}

func NewInterceptor(opt Options) grpc.UnaryServerInterceptor {
	var shouldFire func(method string) (bool, error)
	if !opt.Preprocess {
		shouldFire = lazy
	} else {
		preprocess()
		shouldFire = eager
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			return
		}

		v, err := shouldFire(info.FullMethod)
		if err != nil {
			opt.ErrorLogger.Log(ctx, &Event{
				Id:       "evt_" + slugid.Nice(),
				Type:     info.FullMethod,
				Request:  req.(protoreflect.ProtoMessage),
				Response: resp.(protoreflect.ProtoMessage),
			}, err)
		}

		if !v {
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

		if err := opt.Dispatcher.Dispatch(ctx, evt); err != nil {
			opt.ErrorLogger.Log(ctx, evt, err)
		}

		return
	}
}

func lazy(method string) (bool, error) {
	mutex.RLock()
	v, ok := cache[method]
	mutex.RUnlock()
	if !ok {
		parts := strings.Split(method, "/")
		serviceName := protoreflect.FullName(parts[1])
		methodName := protoreflect.Name(parts[2])

		descriptor, err := protoregistry.GlobalFiles.FindDescriptorByName(serviceName)
		if err != nil {
			return false, err
		}

		v = getFromMethod(descriptor.(protoreflect.ServiceDescriptor).Methods().ByName(methodName))

		mutex.Lock()
		cache[method] = v
		mutex.Unlock()
	}

	return v, nil
}

func preprocess() {
	protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		services := fd.Services()
		for i := 0; i < services.Len(); i++ {
			s := services.Get(i)
			methods := s.Methods()
			for j := 0; j < methods.Len(); j++ {
				m := methods.Get(j)
				cache["/"+string(s.FullName())+"/"+string(m.Name())] = getFromMethod(m)
			}
		}

		return true
	})
}

func getFromMethod(md protoreflect.MethodDescriptor) bool {
	return proto.GetExtension(md.Options(), evtpb.E_Fire).(bool)
}

func eager(method string) (bool, error) {
	return cache[method], nil
}

var cache = map[string]bool{}

var mutex sync.RWMutex

var evtBuf = sync.Pool{
	New: func() interface{} {
		return new(Event)
	},
}
