package main

import (
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
)
import "github.com/sakirsensoy/genv"

var WALLET = ""

func init() {
	WALLET = genv.Key("WALLET").String()
}
