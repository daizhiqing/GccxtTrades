package gateio

import (
	"encoding/json"
	"strconv"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type subModel struct {
	Id     int      `json:"id"`
	Method string   `json:"method"`
	Params []string `json:"params"`
}

type TradeDetail struct {
	tradeList []tradeEntry
	Symbole   string
}

type tradeEntry struct {
	Time    int    `json:"time"`
	Price   string `json:"price"`
	Amount  string `json:"amount"`
	Type    string `json:"type"`
	Symbole string `json:"symbole"`
}

type tradeData struct {
	Method string
	Params []interface{}
}

func GateioWsConnect(sysList []string) {

	if len(sysList) <= 0 {
		logrus.Error("Gateio订阅的交易对数量为空")
		return
	}
	ws, err := websocket.Dial(GateioWsUrl, "", GateioWsUrl)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	subModel := subModel{12312, "trades.subscribe", []string{"BTC_USDT", "EOS_USDT", "EOS_BTC"}}
	subData, err := json.Marshal(subModel)
	if err != nil {
		logrus.Panic("Gateio订阅JSON转换失败")
	}
	logrus.Infof("订阅 %s", subData)
	ws.Write(subData)

	//统计连续错误次数
	var readErrCount = 0

	var msg = make([]byte, GateioBufferSize)

	for {
		if readErrCount > GeteioErrorLimit {
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

		logrus.Infof("Gateio接收：%s \n", msg[:m])
		var t tradeData
		err = json.Unmarshal(msg[:m], &t)
		if err != nil {
			logrus.Error(err)
			continue
		}
		// logrus.Info(t)
		if len(t.Params) > 0 {
			sym := t.Params[0].(string)

			var tradeDetial = TradeDetail{}
			tradeDetial.Symbole = sym
			params := t.Params[1].([]interface{})
			for _, p := range params {
				mapData := p.(map[string]interface{})
				tradeEntry := tradeEntry{
					Time:    int(mapData["time"].(float64)),
					Price:   mapData["price"].(string),
					Amount:  mapData["amount"].(string),
					Type:    mapData["type"].(string),
					Symbole: sym,
				}
				tradeDetial.tradeList = append(tradeDetial.tradeList, tradeEntry)
			}

			logrus.Info("Gateio对象输出", tradeDetial)
		}
	}
}
