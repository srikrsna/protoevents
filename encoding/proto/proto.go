package proto

import (
	"github.com/srikrsna/protoevents"
	"github.com/srikrsna/protoevents/encoding"
	pb "github.com/srikrsna/protoevents/encoding/proto/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var _ encoding.EventMarshaler = (*Marshaler)(nil)

type Marshaler struct{}

func (*Marshaler) MarshalEvent(evt *protoevents.Event) ([]byte, error) {
	ia, err := anypb.New(evt.Request)
	if err != nil {
		return nil, err
	}

	oa, err := anypb.New(evt.Response)
	if err != nil {
		return nil, err
	}

	p := &pb.EventWrapper{
		Id:       evt.Id,
		Type:     evt.Type,
		Request:  ia,
		Response: oa,
	}

	return proto.Marshal(p)
}

var _ encoding.EventUnmarshaler = (*Unmarshaler)(nil)

type Unmarshaler struct{}

func (*Unmarshaler) UnmarshalEvent(b []byte) (*protoevents.Event, error) {
	var w pb.EventWrapper
	if err := proto.Unmarshal(b, &w); err != nil {
		return nil, err
	}

	req, err := w.Request.UnmarshalNew()
	if err != nil {
		return nil, err
	}

	res, err := w.Response.UnmarshalNew()
	if err != nil {
		return nil, err
	}

	return &protoevents.Event{
		Id:       w.Id,
		Type:     w.Type,
		Request:  req,
		Response: res,
	}, nil
}
