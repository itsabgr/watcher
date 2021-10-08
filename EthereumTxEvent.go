package watcher

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/itsabgr/go-handy"
	"gopkg.in/yaml.v3"
	"math/big"
	"time"
)

type EthereumTxEvent interface {
	Event
	WithTx
	WithSender
}

type ethereumTxEvent struct {
	tx     *types.Transaction
	time   time.Time
	sender *common.Address
	block  *big.Int
}

func (n *ethereumTxEvent) MarshalYAML() []byte {
	b, err := yaml.Marshal(struct {
		To, From                                                      *common.Address
		Event, Net, Tx                                                string
		Timestamp, Amount, Gas, GasPrice, GasFeeCap, GasTipCap, Block uint64
	}{
		Event:     n.Kind(),
		Tx:        n.tx.Hash().Hex(),
		Net:       n.Net(),
		Block:     n.block.Uint64(),
		Timestamp: uint64(n.time.Unix()),
		To:        n.tx.To(),
		From:      n.sender,
		Amount:    n.tx.Value().Uint64(),
		Gas:       n.tx.Gas(),
		GasPrice:  n.tx.GasPrice().Uint64(),
		GasFeeCap: n.tx.GasFeeCap().Uint64(),
		GasTipCap: n.tx.GasTipCap().Uint64(),
	})
	handy.Throw(err)
	return b
}
func (n *ethereumTxEvent) Sender() *common.Address {
	return n.sender
}

func (n *ethereumTxEvent) Tx() *types.Transaction {
	return n.tx
}

func (n *ethereumTxEvent) Net() string {
	return "ethereum"
}

func (n *ethereumTxEvent) Kind() string {
	return "transaction"
}

func NewEthereumTxEvent(tx *types.Transaction, sender *common.Address, block *big.Int) (EthereumTxEvent, error) {
	if tx == nil {
		panic(errors.New("nil tx ptr"))
	}
	o := &ethereumTxEvent{}
	o.tx = tx
	o.time = time.Now()
	o.sender = sender
	o.block = block
	return o, nil
}
