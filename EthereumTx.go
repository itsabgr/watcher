package watcher

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type EthereumTx interface {
	Tx
	Index() uint
	EthereumTx() *types.Transaction
	Gas() *big.Int
	GasPrice() *big.Int
	GasFeeCap() *big.Int
	GasTipCap() *big.Int
}

type ethereumTx struct {
	tx     *types.Transaction
	sender *common.Address
	block  *big.Int
	index  uint
}

func (e *ethereumTx) Index() uint {
	return e.index
}

func (e *ethereumTx) Gas() *big.Int {
	return big.NewInt(int64(e.tx.Gas()))
}

func (e *ethereumTx) GasPrice() *big.Int {
	return e.tx.GasPrice()
}

func (e *ethereumTx) GasFeeCap() *big.Int {
	return e.tx.GasFeeCap()
}

func (e *ethereumTx) GasTipCap() *big.Int {
	return e.tx.GasTipCap()
}

func (e *ethereumTx) EthereumTx() *types.Transaction {
	return e.tx
}

func (e *ethereumTx) Block() *big.Int {
	return e.block
}

func (e *ethereumTx) Sender() []byte {
	return e.sender.Bytes()
}

func (e *ethereumTx) Receiver() []byte {
	return e.tx.To().Bytes()
}

func (e *ethereumTx) ID() []byte {
	return e.tx.Hash().Bytes()
}

func (e *ethereumTx) Amount() *big.Int {
	return e.tx.Value()
}

func (e *ethereumTx) Net() string {
	return "ethereum"
}

func (e *ethereumTx) Kind() string {
	return "ethereum"
}

func NewEthereumTx(tx *types.Transaction, sender *common.Address, block *big.Int, index uint) (EthereumTx, error) {
	if tx == nil {
		panic(errors.New("nil tx ptr"))
	}
	o := &ethereumTx{}
	o.tx = tx
	o.sender = sender
	o.block = block
	o.index = index
	return o, nil
}
