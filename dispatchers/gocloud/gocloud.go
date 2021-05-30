package gocloud

import (
	"context"

	"github.com/srikrsna/protoevents"
	"github.com/srikrsna/protoevents/encoding"
	"gocloud.dev/pubsub"
)

var _ protoevents.Dispatcher = (*Dispatcher)(nil)

type Dispatcher struct {
	topic    *pubsub.Topic
	encoding encoding.EventMarshaler
}

func NewDispatcher(topic *pubsub.Topic) *Dispatcher {
	return &Dispatcher{
		topic: topic,
	}
}

func (d *Dispatcher) Dispatch(ctx context.Context, evt *protoevents.Event) error {
	body, err := d.encoding.MarshalEvent(evt)
	if err != nil {
		return err
	}

	m := &pubsub.Message{
		LoggableID: evt.Id,
		Metadata: map[string]string{
			"id":   evt.Id,
			"type": evt.Type,
		},
		Body: body,
	}

	return d.topic.Send(ctx, m)
}
