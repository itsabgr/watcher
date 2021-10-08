package watcher

import (
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)

type EthereumFinEvent interface {
	Event
	Tx() *types.Transaction
}

type ethereumFinEvent struct {
	tx   *types.Transaction
	time time.Time
}

func (n *ethereumFinEvent) Tx() *types.Transaction {
	return n.tx
}

func (n *ethereumFinEvent) Net() string {
	return "ethereum"
}

func (n *ethereumFinEvent) Kind() string {
	return "fin"
}

func NewEthereumFinEvent(tx *types.Transaction) (EthereumFinEvent, error) {
	o := &ethereumFinEvent{}
	o.tx = tx
	o.time = time.Now()
	return o, nil
}
