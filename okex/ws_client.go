package okex

import (
	"encoding/json"
	"errors"

	log "github.com/sirupsen/logrus"

	"strconv"
	"time"

	"golang.org/x/net/websocket"
	"GccxtTrades/config"
	"GccxtTrades/model"
	"strings"
	"GccxtTrades/utils"
)

type tradeDetail struct {
	Channel string     `json:"channel"`
	Data    [][]string `json:"data"`
}

type TradeDetail struct {
	Channel string
	SerNo   string
	Price   string
	Amount  string
	Ts      int64
	Type    string
}

func OkexWsConnect(symbolList []string) {

	if len(symbolList) <= 0 {
		log.Panic(errors.New("Okex订阅的交易对数量为空"))
	}

	id := config.GetExchangeId(Name)
	if id <= 0 {
		log.Println(errors.New(Name + "未找到交易所ID"))
		return
	}
	ws := subWs(symbolList)

	if ws == nil{
		log.Panic("WS连接失败")
	}
	//统计连续错误次数
	var readErrCount = 0
	var msg = make([]byte, OkexBufferSize)

	for {
		if readErrCount > OKexErrLimit {
			//异常退出
			log.Error(errors.New("WebSocket异常连接数连续大于" + strconv.Itoa(readErrCount)))
			ws = subWs(symbolList)
			if ws == nil{
				continue
			}
		}
		m, err := ws.Read(msg)
		if err != nil {
			log.Println(err.Error())
			readErrCount++
			continue
		}
		//连接正常重置
		readErrCount = 0
		log.Printf("Okex接收：%s \n", msg[:m])
		var tradeDetail []tradeDetail
		err = json.Unmarshal(msg[:m], &tradeDetail)
		if err != nil {
			log.Println(err)
			continue
		}
		timeStr := tradeDetail[0].Data[0][3]
		//转成时间戳
		loc, _ := time.LoadLocation("Asia/Shanghai")
		nowStr := time.Now().In(loc).Format("2006-01-02 ")
		tm, err := time.ParseInLocation("2006-01-02 15:04:05", nowStr+timeStr, loc)

		if err != nil {
			log.Println(err)
			continue
		}

		var transData = TradeDetail{
			tradeDetail[0].Channel,
			tradeDetail[0].Data[0][0],
			tradeDetail[0].Data[0][1],
			tradeDetail[0].Data[0][2],
			tm.Unix()*1000,
			tradeDetail[0].Data[0][4]}
		log.Println("Okex输出对象：", transData)

		go DataParser(transData, id)
		go func() {
			select {
			case data := <-model.DataChannel:
				log.Println("获取消息:", data.Symbol, data)
				queueName := config.QueuePre + data.Exchange + "_" + strings.ToLower(strings.Split(data.Symbol, "/")[1])
				utils.SendMsg(config.MqExchange, queueName, data.ToBody())
			default:
				log.Warn(Name + "无消息发送")
			}
		}()
	}

}


func subWs(symbolList []string) *websocket.Conn{
	ws, err := websocket.Dial(OkexWsUrl, "", OkexOrigin)

	if err != nil {
		log.Println(err.Error())
		return nil
	}
	var subList []map[string]string
	//循环订阅交易对
	for _, symbol := range symbolList {

		subStr := map[string]string{"event": "addChannel", "channel": "ok_sub_spot_" + symbol + "_deals"}
		subList = append(subList, subStr)

	}
	sub, _ := json.Marshal(subList)
	_, err = ws.Write(sub)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	log.Printf("订阅: %s \n", sub)
	return  ws
}