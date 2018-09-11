package fcoin

import (
	"encoding/json"
	"strconv"

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
	Ts     int64   `json:"ts"`
	Id     int64   `json:"id"`
	Side   string  `json:"side"`
	Price  float64 `json:"price"`
}

func FcoinWsConnect(symbolList []string) {
	if len(symbolList) <= 0 {
		logrus.Error("Fcoin订阅的交易对数量为空")
		return
	}
	ws, err := websocket.Dial(FCoinWsUrl, "", FCoinWsUrl)
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	for index, s := range symbolList {
		symbolList[index] = "trade." + s
	}
	subModel := subModel{"sub", symbolList}
	subData, err := json.Marshal(subModel)
	if err != nil {
		logrus.Panic("Fcoin订阅JSON转换失败")
	}
	logrus.Infof("订阅 %s", subData)
	ws.Write(subData)

	//统计连续错误次数
	var readErrCount = 0

	var msg = make([]byte, FCoinBufferSize)

	for {
		if readErrCount > FCoinErrorLimit {
			//异常退出
			ws.Close()
			logrus.Panic(("WebSocket异常连接数连续大于" + strconv.Itoa(readErrCount)))

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
	}

}
