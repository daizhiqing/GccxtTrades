package main

import (
	"runtime"
	"ccxt/huobi"
)

func main() {
	runtime.GOMAXPROCS(4)
	list := []string{"eosusdt", "eosbtc", "ethbtc"}
	huobi.HuobiWsConnect(list)
	//zb.ZbWsConnect(list)
}
