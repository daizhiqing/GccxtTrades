package main

import (
	"ccxt/gateio"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

//初始化日志输出格式
func init() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logrus.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	// log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

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
