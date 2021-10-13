package main

import (
	"encoding/hex"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
)
import "github.com/sakirsensoy/genv"

var WALLET []byte = nil

func init() {
	var err error
	WALLET, err = hex.DecodeString(genv.Key("WALLET").String())
	panicerr(err)
}
