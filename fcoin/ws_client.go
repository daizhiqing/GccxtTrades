package fcoin

import (
	"ccxt/config"
	"ccxt/model"
	"ccxt/utils"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type subModel struct {
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

//{"amount":890.000000000,
//	"type":"trade.mitxeth",
//	"ts":1536647971928,
//	"id":310930000,
//	"side":"buy",
//	"price":0.000026940}
type TradeDetail struct {
	//amount
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
	Ts     int     `json:"ts"`
	Id     int64   `json:"id"`
	Side   string  `json:"side"`
	Price  float64 `json:"price"`
}

func FcoinWsConnect(symbolList []string) {
	if len(symbolList) <= 0 {
		logrus.Error("Fcoin订阅的交易对数量为空")
		return
	}
	id := config.GetExchangeId(Name)
	if id <= 0 {
		log.Println(errors.New(Name + "未找到交易所ID"))
		return
	}
	ws := subWs(symbolList)
	if ws == nil {
		logrus.Panic("WS连接失败")
	}
	//统计连续错误次数
	var readErrCount = 0

	var msg = make([]byte, FCoinBufferSize)

	for {
		if readErrCount > FCoinErrorLimit {
			//异常退出
			ws.Close()
			logrus.Error(("WebSocket异常连接数连续大于" + strconv.Itoa(readErrCount)))
			ws = subWs(symbolList)
			if ws == nil{
				continue
			}
		}
		m, err := ws.Read(msg)
		if err != nil {
			logrus.Info(err.Error())
			readErrCount++
			continue
		}
		//连接正常重置
		readErrCount = 0

		logrus.Infof("FCoin接收：%s \n", msg[:m])
		var t TradeDetail
		err = json.Unmarshal(msg[:m], &t)
		if err != nil {
			logrus.Error(err)
			continue
		}
		b, _ := json.Marshal(t)
		logrus.Infoln("FCoin对象输出：", t, string(b))

		go DataParser(t, id)
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

func subWs(symbolList []string) *websocket.Conn {
	ws, err := websocket.Dial(FCoinWsUrl, "", FCoinWsUrl)
	if err != nil {
		logrus.Error(err.Error())
		return nil
	}

	for index, s := range symbolList {
		symbolList[index] = "trade." + s
	}
	subModel := subModel{"sub", symbolList}
	subData, err := json.Marshal(subModel)
	if err != nil {
		logrus.Panic("Fcoin订阅JSON转换失败")
		return nil
	}
	logrus.Infof("订阅 %s", subData)
	ws.Write(subData)
	return ws

}
