package main

import (
	"ccxt/bitfinex"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	bitfinex.StartWs("", false)

	for {
		time.Sleep(time.Hour * 1)
	}
}
