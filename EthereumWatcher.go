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
	lastBlockNumber *big.Int
}

func (e *ethereumWatcher) Pull(ctx context.Context) ([]Event, error) {
again:
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	lastBlockNumber, err := e.client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	lastBlockNumberBig := big.NewInt(int64(lastBlockNumber))
	if e.lastBlockNumber.Cmp(lastBlockNumberBig) == 0 {
		runtime.Gosched()
		goto again
	}
	//
	lastBlock, err := e.client.BlockByNumber(ctx, lastBlockNumberBig)
	if err != nil {
		return nil, err
	}
	lastFinBlock, err := e.client.BlockByNumber(ctx, big.NewInt(0).Sub(lastBlockNumberBig, big.NewInt(6)))
	if err != nil {
		return nil, err
	}
	lastTxs := lastBlock.Transactions()
	lastFinTx := lastFinBlock.Transactions()
	events := make([]Event, 0, lastTxs.Len()+lastFinTx.Len())
	for _, tx := range lastTxs {
		event, err := NewEthereumTxEvent(tx)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	//

	for _, tx := range lastFinTx {
		event, err := NewEthereumFinEvent(tx)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	e.lastBlockNumber = lastBlockNumberBig
	return events, nil
}

func (e *ethereumWatcher) Close() error {
	e.client.Close()
	return nil
}

func NewEthereumWatcher(client *ethclient.Client, lastBlockNum *big.Int) (EthereumWatcher, error) {
	o := &ethereumWatcher{}
	o.client = client
	o.lastBlockNumber = big.NewInt(0).Set(lastBlockNum)
	return o, nil
}
