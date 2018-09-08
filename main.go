package main

import (
	"ccxt/lbank"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	lbank.StartWs("", false)

	for {
		time.Sleep(time.Hour * 1)
	}
}
