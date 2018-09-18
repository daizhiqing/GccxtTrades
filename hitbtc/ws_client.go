package hitbtc

import (
	"GccxtTrades/config"
	"GccxtTrades/model"
	"GccxtTrades/utils"
	"errors"
	"log"
	"strings"

	"encoding/json"
	"strconv"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type tradeInfo struct {
	//{"id":360841011,"price":"0.031089","quantity":"0.977","side":"sell","timestamp":"2018-09-10T12:42:12.905Z"}
	Price     string `json:"price"`
	Quantity  string `json:"quantity"`
	Side      string `json:"side"` //buy sell
	Timestamp string `json:"timestamp"`
}

type paramsEntry struct {
	Data   []tradeInfo `json:"data"`
	Symbol string      `json:"symbol"`
}

type TradeDetail struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  paramsEntry `json:"params"`
}

func HitbtcWsConnect(symbolList []string) {

	if len(symbolList) <= 0 {
		logrus.Error("订阅交易对数量为空")
		return
	}
	id := config.GetExchangeId(Name)
	if id <= 0 {
		log.Println(errors.New(Name + "未找到交易所ID"))
		return
	}
	ws := subWs(symbolList)
	if ws == nil {
		return
	}
	//统计连续错误次数
	var readErrCount = 0

	var msg = make([]byte, HitbtcBufferSize)

	for {
		var data string
		for {
			if readErrCount > HitbtcErrorLimt {
				//异常退出
				ws.Close()
				logrus.Error(errors.New("WebSocket异常连接数连续大于" + strconv.Itoa(readErrCount)))
				ws = subWs(symbolList)
				if ws == nil{
					continue
				}
			}
			m, err := ws.Read(msg)
			if err != nil {
				logrus.Error(err.Error())
				readErrCount++
				continue
			}
			data += string(msg[:m])
			if m <= (HitbtcBufferSize - 1) {
				break
			}
		}
		//连接正常重置
		readErrCount = 0

		logrus.Infof("Hitbtc接收：%s \n", data)
		var t TradeDetail
		err := json.Unmarshal([]byte(data), &t)
		if err != nil {
			logrus.Errorln(err)
			continue
		}
		// logrus.Info("Hitbtc对象输出", t)

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
	//重新订阅
	ws, err := websocket.Dial(HitbtcWsUrl, "", HitbtcWsUrl)
	if err != nil {
		logrus.Error(err.Error())
		return nil
	}

	for _, s := range symbolList {

		subStr := "{\"method\": \"subscribeTrades\", \"params\":{\"symbol\": \"" + s + "\"} ,\"id\": 123}"

		_, err = ws.Write([]byte(subStr))
		if err != nil {
			logrus.Error(err.Error())
			return nil
		}
		logrus.Infof("订阅: %s \n", subStr)
	}
	return ws
}
