package huobi

import (
	"bytes"
	"ccxt/config"
	"compress/gzip"
	"encoding/json"
	"errors"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"log"
	"math/big"
	"strconv"
	"strings"
)

type subModel struct {
	Sub string `json:"sub"`
	Id  int    `json:"id"`
}

type trade struct {
	Id        big.Int `json:"id"`
	Price     float32 `json:"price"`
	Direction string  `json:"direction"`
	Amount    float32 `json:"amount"`
	Ts        int     `json:"ts"`
}

type tick struct {
	Id   int     `json:"id"`
	Ts   int     `json:"ts"`
	Data []trade `json:"data"`
}

type TradeDetail struct {
	Ch   string `json:"ch"`
	Ts   int    `json:"ts"`
	Tick tick   `json:"tick"`
}

func HuobiWsConnect(symbolList []string) {

	if len(symbolList) <= 0 {
		log.Panic(errors.New("火币订阅的交易对数量为空"))
	}

	ws, err := websocket.Dial(config.HuoBiWsUrl, "", config.HuoBiOrigin)

	if err != nil {
		log.Println(err.Error())
		return
	}
	//循环订阅交易对
	for _, symbol := range symbolList {
		sub := subModel{"market." + symbol + ".trade.detail", config.HuoBiGId}
		message, err := json.Marshal(sub)
		if err != nil {
			log.Println(err.Error())
			return
		}
		_, err = ws.Write(message)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Printf("订阅: %s \n", message)
	}
	//统计连续错误次数
	var readErrCount = 0
	var msg = make([]byte, config.HuoBiMsgBufferSize)
	for {
		if readErrCount > config.HuoBiErroLimit {
			//异常退出
			log.Panic(errors.New("WebSocket异常连接数连续大于" + strconv.Itoa(readErrCount)))
		}
		m, err := ws.Read(msg)
		if err != nil {
			log.Println(err.Error())
			readErrCount++
			continue
		}
		//连接正常重置
		readErrCount = 0
		reader, err := gzip.NewReader(bytes.NewReader(msg[:m]))
		if err != nil {
			log.Println(err)
			continue
		}
		b, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Println(err)
			continue
		}
		revMsg := string(b)
		//ping pong 心跳防止断开
		if strings.Contains(revMsg, "ping") {
			ws.Write([]byte(strings.Replace(revMsg, "ping", "pong", 1)))
		}
		log.Println("Huobi接收：", revMsg)
		var tradeDetail TradeDetail
		err = json.Unmarshal(b, &tradeDetail)
		if err != nil {
			log.Println(err)
			continue
		}
		//json , _ :=simplejson.NewJson(b)
		//temp ,_ :=json.Marshal(tradeDetail)
		if tradeDetail.Ch != "" {
			//log.Println("转化：", string(temp))
			log.Println("Huobi输出对象：", tradeDetail)
		}
	}
	ws.Close() //关闭连接
}
