package watcher

import "math/big"

type EthereumBlock interface {
	Txs() []EthereumTx
	Number() *big.Int
}
