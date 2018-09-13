package main

import (
	"ccxt/gateio"
	"runtime"

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

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			logrus.Error(err) // 这里的err其实就是panic传入的内容
		}
		logrus.Error("程序进程退出")
	}()

	runtime.GOMAXPROCS(runtime.NumCPU())

	// bitfinex.StartWs("", false)
	// huobi.StartWs("", false)
	// lbank.StartWs("", false)
	// okex.StartWs("", false)
	// binance.StartWs("", false)
	// zb.StartWs("", false)
	gateio.StartWs("", false)
	// hitbtc.StartWs("", false)
	// fcoin.StartWs("", false)
	// hadax.StartWs("", false)
	// go func() {
	// 	for {
	// 		utils.SendMsg("ex-api-mq", "trades_binance_btc", []byte("go-1:"+time.Now().String()))
	// 		// time.Sleep(time.Second * 1)
	// 	}
	// }()

	// go utils.ReceiveMsg("goDzq", "trades_binance_btc", func(b []byte) {
	// 	logrus.Errorf("trades_binance_btc : %s", b)
	// })
	loop := make(chan bool)
	<-loop
}
