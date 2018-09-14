package zb

import (
	"encoding/json"
	"errors"

	log "github.com/sirupsen/logrus"

	"strconv"

	"golang.org/x/net/websocket"
	"ccxt/config"
)

type detail struct {
	Amount     string `json:"amount"`
	Price      string `json:"price"`
	Tid        int    `json:"tid"`
	Date       int    `json:"date"`
	Type       string `json:"type"` // sell  || buy
	Trade_type string `json:"trade_type"`
}

type TradeDetail struct {
	DataType string   `json:"dataType"`
	Data     []detail `json:"data"`
	Channel  string   `json:"channel"`
}

func ZbWsConnect(symbolList []string) {

	if len(symbolList) <= 0 {
		log.Println(errors.New("ZB订阅的交易对数量为空"))
		return
	}
	id := config.GetExchangeId(Name)
	if id <= 0 {
		log.Println(errors.New(Name + "未找到交易所ID"))
		return
	}
	ws := subWs(symbolList)

	//统计连续错误次数
	var readErrCount = 0

	var msg = make([]byte, ZbMsgBufferSize)

	for {
		var data string
		for {
			if readErrCount > ZbErrorLimit {
				//异常退出
				ws.Close()
				log.Error(errors.New("WebSocket异常连接数连续大于" + strconv.Itoa(readErrCount)))
				ws = subWs(symbolList)
			}
			m, err := ws.Read(msg)
			if err != nil {
				log.Error(err.Error())
				readErrCount++
				continue
			}
			data += string(msg[:m])
			if m <= (ZbMsgBufferSize - 1) {
				break
			}
		}
		//连接正常重置
		readErrCount = 0
		log.Printf("Zb接收：%s \n", data)
		var tradeDetail TradeDetail
		err := json.Unmarshal([]byte(data), &tradeDetail)
		if err != nil {
			log.Println(err)
			continue
		}
		if tradeDetail.Channel != "" {
			log.Println("Zb输出对象：", tradeDetail)
		}
	}
}

func subWs(symbolList []string) *websocket.Conn {
	ws, err := websocket.Dial(ZbWsUrl, "", ZbOrigi)

	if err != nil {
		log.Println(err.Error())
		return nil
	}
	//循环订阅交易对
	for _, symbol := range symbolList {

		subStr := "{\"event\":\"addChannel\",\"channel\":\"" + symbol + "_trades\"}"
		//subStr := "{\"event\":\"addChannel\",\"channel\":\"btcusdt_trades\"}"
		_, err = ws.Write([]byte(subStr))
		if err != nil {
			log.Println(err.Error())
			return nil
		}
		log.Printf("订阅: %s \n", subStr)
	}

	return ws
}