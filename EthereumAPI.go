package watcher

import (
	"context"

	"github.com/ybbus/jsonrpc"
	"io"
	"math/big"
	"os"
)

var ErrNotFound = os.ErrNotExist

type EthereumAPIClient interface {
	GetTxByBlockNumberAndIndex(ctx context.Context, block *big.Int, index uint) (EthereumTx, error)
	GetBlockByNumber(context.Context, *big.Int) (EthereumBlock, error)
	GetLastBlockNumber(context.Context) (*big.Int, error)
	io.Closer
}

type ethereumInfuraClient struct {
	url string
	cli *jsonrpc.RPCClient
}

func (e *ethereumInfuraClient) GetTxByBlockNumberAndIndex(ctx context.Context, block *big.Int, index uint) (EthereumTx, error) {

}

func (e *ethereumInfuraClient) GetBlockByNumber(ctx context.Context, b *big.Int) (EthereumBlock, error) {
}

func (e *ethereumInfuraClient) GetLastBlockNumber(ctx context.Context) (*big.Int, error) {

}

func (e *ethereumInfuraClient) Close() error {
	return nil
}

func NewEthereumInfuraClient(url string) (EthereumAPIClient, error) {
	cli := new(ethereumInfuraClient)
	cli.cli = jsonrpc.NewRPCClient(url)
	cli.url = url
	return cli, nil
}
