package watcher

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/itsabgr/go-handy"
)

func ethereumTxToString(tx *types.Transaction) string {
	b, err := tx.MarshalJSON()
	handy.Throw(err)
	return string(b)
}
