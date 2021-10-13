package watcher

import (
	"errors"
	"github.com/itsabgr/go-handy"
	"gopkg.in/yaml.v3"
	"math/big"
	"time"
)

type EthereumFinEvent interface {
	Event
	WithTx
}

type ethereumFinEvent struct {
	tx    EthereumTx
	time  time.Time
	block *big.Int
}

func (n *ethereumFinEvent) Tx() Tx {
	return n.tx
}

func (n *ethereumFinEvent) Net() string {
	return "ethereum"
}

func (n *ethereumFinEvent) MarshalYAML() []byte {
	b, err := yaml.Marshal(struct {
		To, From, Tx                                                  []byte
		Event, Net                                                    string
		Timestamp, Amount, Gas, GasPrice, GasFeeCap, GasTipCap, Block uint64
	}{
		Event:     n.Kind(),
		Tx:        n.tx.ID(),
		Block:     n.block.Uint64(),
		Net:       n.Net(),
		Timestamp: uint64(n.time.Unix()),
		To:        n.tx.Receiver(),
		From:      n.tx.Sender(),
		Amount:    n.tx.Amount().Uint64(),
		Gas:       n.tx.Gas().Uint64(),
		GasPrice:  n.tx.GasPrice().Uint64(),
		GasFeeCap: n.tx.GasFeeCap().Uint64(),
		GasTipCap: n.tx.GasTipCap().Uint64(),
	})
	handy.Throw(err)
	return b
}
func (n *ethereumFinEvent) Kind() string {
	return "finality"
}

func NewEthereumFinEvent(tx EthereumTx) (EthereumFinEvent, error) {
	if tx == nil {
		panic(errors.New("nil tx ptr"))
	}
	o := &ethereumFinEvent{}
	o.tx = tx
	o.time = time.Now()
	return o, nil
}
