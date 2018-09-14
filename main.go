package main

import (
	"runtime"

	"github.com/sirupsen/logrus"
	"ccxt/binance"
	"flag"
	"ccxt/bitfinex"
	"ccxt/huobi"
	"ccxt/lbank"
	"ccxt/okex"
	"ccxt/zb"
	"ccxt/gateio"
	"ccxt/hitbtc"
	"ccxt/fcoin"
	"ccxt/hadax"
	"ccxt/utils"
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

	name := flag.String("name", "", "交易所名称")
	mq := flag.String("mq", "", "amqp://user:pwd@host:port/vhost")

	flag.Parse()
	logrus.Info(*name , *mq)

	utils.AmqpUrl = *mq
	runtime.GOMAXPROCS(runtime.NumCPU())

	switch *name {
	case binance.Name:
		binance.StartWs("", false)
	case bitfinex.Name:
		bitfinex.StartWs("", false)
	case huobi.Name:
		huobi.StartWs("", false)
	case gateio.Name:
		gateio.StartWs("", false)
	case hitbtc.Name:
		hitbtc.StartWs("", false)
	case fcoin.Name:
		fcoin.StartWs("", false)
	case hadax.Name:
		hadax.StartWs("", false)
	case lbank.Name:
		lbank.StartWs("", false)
	case okex.Name:
		okex.StartWs("", false)
	case zb.Name:
		zb.StartWs("", false)
	default:
		logrus.Panic("name is not set")
	}
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
