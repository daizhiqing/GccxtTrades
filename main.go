package main

import (
	"ccxt/binance"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	binance.StartWs("", false)

	for {
		time.Sleep(time.Hour * 1)
	}
}
