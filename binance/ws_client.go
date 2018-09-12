package binance

import (
	"encoding/json"
	"errors"

	"github.com/sirupsen/logrus"

	log "github.com/sirupsen/logrus"

	"ccxt/config"
	"ccxt/model"
	"ccxt/utils"
	"strconv"
	"strings"

	"golang.org/x/net/websocket"
)

type data struct {
	EventType string `json:"e"`
	EventTime int    `json:"E"`
	Symbol    string `json:"s"`
	TradeId   int    `json:"t"`
	Price     string `json:"p"`
	Quantity  string `json:"q"`
	BuyerId   int    `json:"b"`
	SellerId  int    `json:"a"`
	Ts        int    `json:"T"`
	Buy       bool   `json:"m"`
	M         bool   `json:"M"`
}

type TradeDatail struct {
	Stream string `json:"stream"`
	Data   data   `json:"data"`
}

func BinanceWsConnect(symbolList []string) {
	if len(symbolList) <= 0 {
		log.Println(errors.New("Binance订阅的交易对数量为空"))
		return
	}

	var subUrl string
	for _, s := range symbolList {
		subUrl += strings.ToLower(s) + "@aggTrade/"
	}
	id := config.GetExchangeId(Name)

	if id <= 0 {
		log.Println(errors.New(Name + "未找到交易ID"))
		return
	}

	ws, err := websocket.Dial(BinanceWsUrl+subUrl, "", BinanceOrigin)
	log.Printf("订阅: %s \n", subUrl)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//统计连续错误次数
	var readErrCount = 0
	var msg = make([]byte, BinanceBufferSize)
	for {
		if readErrCount > BinanceErrorLimit {
			ws.Close()
			log.Panic(errors.New("WebSocket异常连接数连续大于" + strconv.Itoa(readErrCount)))
			break
		}
		m, err := ws.Read(msg)
		if err != nil {
			log.Println(err.Error())
			readErrCount++
			continue
		}
		//连接正常重置
		readErrCount = 0

		// log.Printf("Binance接收：%s \n", msg[:m])

		var t TradeDatail
		err = json.Unmarshal(msg[:m], &t)
		if err != nil {
			log.Println(err)
			continue
		}

		// log.Println("Binance输出对象：", t.Data.Buy, t.Data.M, t)

		go DataParser(t, id)
		go func() {
			select {
			case data := <-model.DataChannel:
				queueName := config.QueuePre + data.Exchange + "_" + strings.ToLower(strings.Split(data.Symbol, "/")[1])
				utils.SendMsg(config.MqExchange, queueName, data.ToBody())
			default:
				logrus.Warn("Binance无消息发送")
			}
		}()
	}

}
