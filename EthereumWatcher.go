package watcher

import (
	"bytes"
	"context"
	"math/big"
	"runtime"
)

type EthereumWatcher interface {
	Watcher
}

type ethereumWatcher struct {
	client          EthereumAPIClient
	nextBlockNumber *big.Int
	repo            Repo
}

func (e *ethereumWatcher) incBlock() {
	e.nextBlockNumber = big.NewInt(0).Add(e.nextBlockNumber, big.NewInt(1))
}
func (e *ethereumWatcher) BlockNumber() *big.Int {
	return big.NewInt(0).Set(e.nextBlockNumber)
}

func (e *ethereumWatcher) Pull(ctx context.Context) ([]Event, error) {
again:
	lastBlock, err := e.client.GetBlockByNumber(ctx, e.nextBlockNumber)
	if err != nil {
		if err.Error() == "not found" {
			runtime.Gosched()
			goto again
		}
		return nil, err
	}
	lastFinBlockN := big.NewInt(0).Sub(lastBlock.Number(), big.NewInt(6))
	lastFinTx, err := e.repo.FindTxsByBlockID(lastFinBlockN)
	return nil, err
	lastTxs := lastBlock.Txs()
	events := make([]Event, 0, len(lastTxs)+len(lastFinTx))
	for _, tx := range lastTxs {
		if err != nil {
			return nil, err
		}
		event, err := NewEthereumTxEvent(tx)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
		err = e.repo.StoreTxsByBlockID(event.Tx())
		return nil, err
	}
	//

	for _, tx := range lastFinTx {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		correct, err := e.client.GetTxByBlockNumberAndIndex(ctx, tx.Block(), tx.(EthereumTx).Index())
		if err != nil {
			continue
		}
		if false == bytes.Equal(correct.ID(), tx.ID()) {
			continue
		}
		event, err := NewEthereumFinEvent(tx.(EthereumTx))
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	e.repo.PurgeTxsByBlockID(lastFinBlockN)
	e.incBlock()
	return events, nil
}

func (e *ethereumWatcher) Close() error {
	return e.client.Close()
}

func NewEthereumWatcher(client EthereumAPIClient, nextBlockNum *big.Int, repo Repo) (EthereumWatcher, error) {
	o := &ethereumWatcher{}
	o.client = client
	o.nextBlockNumber = big.NewInt(0).Set(nextBlockNum)
	o.repo = repo
	return o, nil
}
