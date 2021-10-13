package main

import (
	"bytes"
	"github.com/itsabgr/watcher"
	"os"
)

func main() {
	eventLog, err := watcher.NewFileEventLog(os.Stdout)
	panicerr(err)
	defer eventLog.Close()
	ethClient, err := watcher.NewEthereumInfuraClient(NET)
	panicerr(err)
	defer ethClient.Close()
	startBlockNumber, err := ethClient.GetLastBlockNumber(ctx)
	panicerr(err)
	repo, err := watcher.NewMemRepo()
	panicerr(err)
	ethereumWatcher, err := watcher.NewEthereumWatcher(ethClient, startBlockNumber, repo)
	panicerr(err)
	defer ethereumWatcher.Close()
	for {
		events, err := ethereumWatcher.Pull(ctx)
		panicerr(err)
		for _, event := range events {
			if WALLET != nil {
				switch event.(type) {
				case watcher.EthereumTxEvent, watcher.EthereumFinEvent:
					switch {
					case bytes.Equal(event.(watcher.WithTx).Tx().Sender(), WALLET),
						bytes.Equal(event.(watcher.WithTx).Tx().Receiver(), WALLET):
					default:
						continue
					}
				default:
					continue
				}
			}
			panicerr(eventLog.Emit(event))
		}
	}
}
