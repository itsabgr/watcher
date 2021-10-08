package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/itsabgr/watcher"
	"math/big"
	"os"
)

func main() {
	eventLog, err := watcher.NewFileEventLog(os.Stdout)
	panicerr(err)
	defer eventLog.Close()
	ethClient, err := ethclient.DialContext(ctx, NET)
	panicerr(err)
	defer ethClient.Close()
	startBlockNumber, err := ethClient.BlockNumber(ctx)
	panicerr(err)
	ethereumWatcher, err := watcher.NewEthereumWatcher(ethClient, big.NewInt(int64(startBlockNumber)))
	panicerr(err)
	defer ethereumWatcher.Close()
	for {
		events, err := ethereumWatcher.Pull(ctx)
		panicerr(err)
		for _, event := range events {
			if WALLET != "" {
				switch event.(type) {
				case watcher.EthereumTxEvent, watcher.EthereumFinEvent:
					if event.(watcher.WithSender).Sender().Hex() != WALLET {
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
