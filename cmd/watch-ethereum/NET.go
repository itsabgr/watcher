package main

import (
	"errors"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
)
import "github.com/sakirsensoy/genv"

var NET = ""

func init() {
	net := genv.Key("NET").String()
	if net == "" {
		panicerr(errors.New("need NET environment variable"))
	}
	NET = net
}
