package hadax

import (
	"bytes"
	"ccxt/config"
	"ccxt/model"
	"ccxt/utils"
	"compress/gzip"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type subModel struct {
	Sub string `json:"sub"`
	Id  int    `json:"id"`
}

type trade struct {
	// Id        int64   `json:"id"`
	Price     float64 `json:"price"`
	Direction string  `json:"direction"`
	Amount    float64 `json:"amount"`
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

func HadaxWsConnect(symbolList []string) {

	if len(symbolList) <= 0 {
		log.Println(errors.New("Hadax订阅的交易对数量为空"))
		return
	}
	id := config.GetExchangeId(Name)
	if id <= 0 {
		log.Println(errors.New(Name + "未找到交易所ID"))
		return
	}
	ws, err := websocket.Dial(HadaxWsUrl, "", HadaxWsUrl)

	if err != nil {
		log.Println(err.Error())
		return
	}
	//循环订阅交易对
	for _, symbol := range symbolList {
		sub := subModel{"market." + symbol + ".trade.detail", 1001}
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
	var msg = make([]byte, HadaxBufferSize)
	for {
		if readErrCount > HadaxErrorLimit {
			//异常退出
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
		log.Println("Hadax接收：", revMsg)
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
			// log.Println("Hadax输出对象：", tradeDetail)

			go DataParser(tradeDetail, id)
			go func() {
				select {
				case data := <-model.DataChannel:
					log.Println("获取消息:", data.Symbol, data)
					queueName := config.QueuePre + data.Exchange + "_" + strings.ToLower(strings.Split(data.Symbol, "/")[1])
					utils.SendMsg(config.MqExchange, queueName, data.ToBody())
				default:
					logrus.Warn(Name + "无消息发送")
				}
			}()
		}
	}

}
