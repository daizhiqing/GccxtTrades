package main

import (
	"ccxt/huobi"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	list := []string{"eosusdt", "eosbtc", "ethbtc"}
	huobi.HuobiWsConnect(list)
}
