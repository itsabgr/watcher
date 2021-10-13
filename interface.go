package watcher

import (
	"context"
	"io"
	"math/big"
)

type Repo interface {
	StoreTxsByBlockID(Tx) error
	FindTxsByBlockID(blockID *big.Int) ([]Tx, error)
	PurgeTxsByBlockID(blockID *big.Int) error
}
type Tx interface {
	Sender() []byte
	Receiver() []byte
	ID() []byte
	Amount() *big.Int
	Net() string
	Kind() string
	Block() *big.Int
}
type WithTx interface {
	Tx() Tx
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
