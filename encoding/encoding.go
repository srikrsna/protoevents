package encoding

import "github.com/srikrsna/protoevents"

type EventMarshaler interface {
	MarshalEvent(*protoevents.Event) ([]byte, error)
}

type EventUnmarshaler interface {
	UnmarshalEvent([]byte) (*protoevents.Event, error)
}
