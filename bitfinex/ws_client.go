package bitfinex

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"golang.org/x/net/websocket"
)

type TradeDatial struct {
	Symbole string
	Ts      int
	Price   float64
	Amount  float64
}

func BitfinexWsConnect(symbolList []string) {
	if len(symbolList) <= 0 {
		log.Println(errors.New("Binance订阅的交易对数量为空"))
		return
	}
	ws, err := websocket.Dial(BitfinexWsUrl, "", BitfinexOrigin)
	if err != nil {
		log.Println(err.Error())
		return
	}
	for _, s := range symbolList {
		subStr := "{\"event\": \"subscribe\", \"channel\": \"trades\", \"pair\":\"" + s + "\" }"

		_, err = ws.Write([]byte(subStr))
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Printf("订阅: %s \n", subStr)
	}

	//统计连续错误次数
	var readErrCount = 0

	var msg = make([]byte, BitfinexBufferSize)

	for {
		if readErrCount > BitfinexErrorLimit {
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

		log.Printf("Bitfinex接收：%s \n", msg[:m])

		var revData = make([]interface{}, 7)
		err = json.Unmarshal(msg[:m], &revData)
		if err != nil {
			log.Println(err)
			continue
		}
		// revDataStr := string(msg[:m])
		// revDataStr = strings.Replace(revDataStr , "[" ,-1)
		// revDataStr = strings.Replace(revDataStr , "]" ,-1)
		// revDataStr = strings.Replace(revDataStr , "]" ,-1)
		// revDataStr = strings.Replace(revDataStr , "\"" ,-1)

		// revData := strings.Split(revDataStr , ",")
		if revData[1] == "tu" {
			t := TradeDatial{
				revData[2].(string),
				int(revData[4].(float64)),
				revData[5].(float64),
				revData[6].(float64)}
			log.Println("Bitfinex输出对象：", t)
		}
	}
}
