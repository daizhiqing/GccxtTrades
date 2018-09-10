package main

import (
	"ccxt/gateio"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)

	// bitfinex.StartWs("", false)
	// huobi.StartWs("", false)
	// lbank.StartWs("", false)
	// okex.StartWs("", false)
	// binance.StartWs("", false)
	// zb.StartWs("", false)
	gateio.StartWs("", false)

	for {
		time.Sleep(time.Hour * 1)
	}
}
