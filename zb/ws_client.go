package zb

import (
	"ccxt/config"
	"errors"
	"golang.org/x/net/websocket"
	"log"
	"strconv"
	"encoding/json"
)

type detail struct {
	Amount string `json:"amount"`
	Price string `json:"price"`
	Tid int `json:"tid"`
	Date int `json:"date"`
	Type string `json:"type"`  // sell  || buy
	Trade_type string `json:"trade_type"`
}

type TradeDetail struct {
	DataType string `json:"dataType"`
	Data []detail `json:"data"`
	Channel string `json:"channel"`
}

func ZbWsConnect(symbolList []string) {

	if len(symbolList) <= 0 {
		log.Println(errors.New("ZB订阅的交易对数量为空"))
		return
	}

	ws, err := websocket.Dial(config.ZbWsUrl, "", config.ZbOrigi)

	if err != nil {
		log.Println(err.Error())
		return
	}
	//循环订阅交易对
	for _, symbol := range symbolList {

		subStr := "{\"event\":\"addChannel\",\"channel\":\"" + symbol + "_trades\"}"
		//subStr := "{\"event\":\"addChannel\",\"channel\":\"btcusdt_trades\"}"
		_, err = ws.Write([]byte(subStr))
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Printf("订阅: %s \n", subStr)
	}

	//统计连续错误次数
	var readErrCount = 0

	var msg = make([]byte, config.ZbMsgBufferSize)

	for {
		if readErrCount > config.ZbErrorLimit {
			//异常退出
			log.Panic(errors.New("WebSocket异常连接数连续大于" + strconv.Itoa(readErrCount)))
		}
		m, err := ws.Read(msg)
		if err != nil {
			log.Println(err.Error())
			readErrCount ++
			continue
		}
		//连接正常重置
		readErrCount = 0
		log.Printf("Zb接收：%s \n", msg[:m])
		var tradeDetail TradeDetail
		err = json.Unmarshal(msg[:m], &tradeDetail)
		if err != nil {
			log.Println(err)
			continue
		}
		if tradeDetail.Channel != "" {
			log.Println("Zb输出对象：", tradeDetail)
		}
	}
	ws.Close() //关闭连接
}
