package watcher

import (
	"errors"
	"github.com/itsabgr/go-handy"
	"gopkg.in/yaml.v3"
	"math/big"
	"time"
)

type EthereumTxEvent interface {
	Event
	WithTx
}

type ethereumTxEvent struct {
	tx    EthereumTx
	time  time.Time
	block *big.Int
}

func (n *ethereumTxEvent) MarshalYAML() []byte {
	b, err := yaml.Marshal(struct {
		To, From, Tx                                                  []byte
		Event, Net                                                    string
		Timestamp, Amount, Gas, GasPrice, GasFeeCap, GasTipCap, Block uint64
	}{
		Event:     n.Kind(),
		Tx:        n.tx.ID(),
		Net:       n.Net(),
		Block:     n.block.Uint64(),
		Timestamp: uint64(n.time.Unix()),
		To:        n.tx.Receiver(),
		From:      n.tx.Receiver(),
		Amount:    n.tx.Amount().Uint64(),
		Gas:       n.tx.Gas().Uint64(),
		GasPrice:  n.tx.GasPrice().Uint64(),
		GasFeeCap: n.tx.GasFeeCap().Uint64(),
		GasTipCap: n.tx.GasTipCap().Uint64(),
	})
	handy.Throw(err)
	return b
}

func (n *ethereumTxEvent) Tx() Tx {
	return n.tx
}

func (n *ethereumTxEvent) Net() string {
	return "ethereum"
}

func (n *ethereumTxEvent) Kind() string {
	return "transaction"
}

func NewEthereumTxEvent(tx EthereumTx) (EthereumTxEvent, error) {
	if tx == nil {
		panic(errors.New("nil tx ptr"))
	}
	o := &ethereumTxEvent{}
	o.tx = tx
	o.time = time.Now()
	return o, nil
}
