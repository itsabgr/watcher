package watcher

import (
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)

type EthereumTxEvent interface {
	Event
	Tx() *types.Transaction
}

type ethereumTxEvent struct {
	tx   *types.Transaction
	time time.Time
}

func (n *ethereumTxEvent) Tx() *types.Transaction {
	return n.tx
}

func (n *ethereumTxEvent) Net() string {
	return "ethereum"
}

func (n *ethereumTxEvent) Kind() string {
	return "tx"
}


func NewEthereumTxEvent(tx *types.Transaction) (EthereumTxEvent, error) {
	o := &ethereumTxEvent{}
	o.tx = tx
	o.time = time.Now()
	return o, nil
}
