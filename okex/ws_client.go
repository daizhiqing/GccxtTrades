package okex

import (
	"log"
	"golang.org/x/net/websocket"
	"ccxt/config"
	"errors"
	"strconv"
)

func OkexWsConnect(symbolList []string) {

	if len(symbolList) <= 0 {
		log.Panic(errors.New("Okex订阅的交易对数量为空"))
	}

	ws, err := websocket.Dial(config.OkexWsUrl, "", config.OkexOrigin)

	if err != nil {
		log.Println(err.Error())
		return
	}
	//循环订阅交易对
	for _, symbol := range symbolList {

		subStr := "{'event':'addChannel','channel':'ok_sub_spot_" + symbol + "_deals'}"
		_, err = ws.Write([]byte(subStr))
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Printf("订阅: %s \n", subStr)
	}

	//统计连续错误次数
	var readErrCount= 0
	var msg= make([]byte, config.OkexBufferSize)

	for {
		if readErrCount > config.ZbErrorLimit {
			//异常退出
			log.Panic(errors.New("WebSocket异常连接数连续大于" + strconv.Itoa(readErrCount)))
		}
		m, err := ws.Read(msg)
		if err != nil {
			log.Println(err.Error())
			readErrCount ++
			continue
		}
		//连接正常重置
		readErrCount = 0
		log.Printf("Okex接收：%s \n", msg[:m])
	}
	ws.Close() //关闭连接
}