package main

import (
	"runtime"
	"time"
	"ccxt/bitfinex"
)

func main() {
	runtime.GOMAXPROCS(4)
	bitfinex.StartWs("", false)

	for {
		time.Sleep(time.Hour * 1)
	}
}
