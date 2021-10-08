package watcher

import (
	"context"
	"io"
)

type EventLog interface {
	Emit(Event) error
	io.Closer
}

type Event interface {
	Kind() string
	Net() string
}

type Watcher interface {
	Pull(ctx context.Context) ([]Event, error)
	io.Closer
}
type FinEvent interface {
	Event
	TxID() []byte
}
