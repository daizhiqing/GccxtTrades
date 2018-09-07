package main

import (
	"runtime"
	"time"
	"ccxt/utils"
	"ccxt/config"
	"strings"
	"log"
	"ccxt/zb"
	"ccxt/huobi"
)

func main() {
	runtime.GOMAXPROCS(4)
	//获取火币的所以交易对
	huobiSymbols := utils.HttpGet(config.HuoBiSymbols).Get("data").MustArray()
	var syList []string
	for _,m := range huobiSymbols {
		str := m.(map[string]interface{})["base-currency"].(string)+ m.(map[string]interface{})["quote-currency"].(string)
		syList = append(syList , str)
	}
	log.Println("huobi:" , syList)

	go huobi.HuobiWsConnect(syList)

	var syListZb []string
	zbSym,_ := utils.HttpGet(config.ZbSymbols).Map()
	for key ,_ := range zbSym {
		syListZb = append(syListZb , strings.Replace(key , "_" , "" , -1))
	}
	log.Println("ZB:" , syListZb)
	go zb.ZbWsConnect(syListZb)
	for  {
		time.Sleep(time.Second * 50)
	}
}
