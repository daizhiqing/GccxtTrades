package main

import (
	"runtime"
	"ccxt/okex"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	okex.StartWs("" , true)

	for  {
		time.Sleep(time.Hour*1)
	}
}
