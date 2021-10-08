package watcher

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"io"
)

type WithTx interface {
	Tx() *types.Transaction
}
type WithSender interface {
	Sender() *common.Address
}
type EventLog interface {
	Emit(Event) error
	io.Closer
}

type Event interface {
	Kind() string
	Net() string
	MarshalYAML() []byte
}

type Watcher interface {
	Pull(ctx context.Context) ([]Event, error)
	io.Closer
}
type FinEvent interface {
	Event
	TxID() []byte
}
