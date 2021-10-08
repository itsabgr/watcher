package main

import (
	"os"
	"os/signal"
)

func init() {
	var sig = make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	go func() {
		panicerr(<-sig)
	}()
}
