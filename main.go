package main

import (
	"runtime"
	"time"
	"ccxt/lbank"
)

func main() {
	runtime.GOMAXPROCS(4)
	lbank.StartWs("" , false)

	for  {
		time.Sleep(time.Hour*1)
	}
}
