package watcher

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"runtime"
)

type EthereumWatcher interface {
	Watcher
}

type ethereumWatcher struct {
	client          *ethclient.Client
	nextBlockNumber *big.Int
}

func (e *ethereumWatcher) incBlock() {
	e.nextBlockNumber = big.NewInt(0).Add(e.nextBlockNumber, big.NewInt(1))
}
func (e *ethereumWatcher) BlockNumber() *big.Int {
	return big.NewInt(0).Set(e.nextBlockNumber)
}

func (e *ethereumWatcher) Pull(ctx context.Context) ([]Event, error) {
again:
	lastBlock, err := e.client.BlockByNumber(ctx, e.nextBlockNumber)
	if err != nil {
		if err.Error() == "not found" {
			runtime.Gosched()
			goto again
		}
		return nil, err
	}
	lastFinBlock, err := e.client.BlockByNumber(ctx, big.NewInt(0).Sub(e.nextBlockNumber, big.NewInt(6)))
	if err != nil {
		return nil, err
	}
	lastTxs := lastBlock.Transactions()
	lastFinTx := lastFinBlock.Transactions()
	events := make([]Event, 0, lastTxs.Len()+lastFinTx.Len())
	for i, tx := range lastTxs {
		sender, err := e.client.TransactionSender(ctx, tx, lastBlock.Hash(), uint(i))
		if err != nil {
			return nil, err
		}
		event, err := NewEthereumTxEvent(tx, &sender, lastBlock.Number())
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	//

	for i, tx := range lastFinTx {
		sender, err := e.client.TransactionSender(ctx, tx, lastFinBlock.Hash(), uint(i))
		if err != nil {
			return nil, err
		}
		event, err := NewEthereumFinEvent(tx, &sender, lastFinBlock.Number())
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	e.incBlock()
	return events, nil
}

func (e *ethereumWatcher) Close() error {
	e.client.Close()
	return nil
}

func NewEthereumWatcher(client *ethclient.Client, nextBlockNum *big.Int) (EthereumWatcher, error) {
	o := &ethereumWatcher{}
	o.client = client
	o.nextBlockNumber = big.NewInt(0).Set(nextBlockNum)
	return o, nil
}
