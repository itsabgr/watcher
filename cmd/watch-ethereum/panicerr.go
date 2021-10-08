package main

import (
	"context"
	"log"
	"os"
)

var ctx, cancel = context.WithCancel(context.Background())

func panicerr(err interface{}) {
	if err == nil {
		return
	}
	cancel()
	log.Println(err)
	os.Exit(1)
}
