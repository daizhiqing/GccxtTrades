package lbank

import (
	"encoding/json"
	"errors"

	log "github.com/sirupsen/logrus"

	"strconv"
	"strings"

	"golang.org/x/net/websocket"
)

type trade struct {
	Volume    float32 `json:"volume"`
	Price     float32 `json:"price"`
	Amount    float32 `json:"amount"`
	Direction string  `json:"direction"`
	TS        string  `json:"TS"`
}

type TradeDetail struct {
	Pair   string `json:"pair"`
	Trade  trade  `json:"trade"`
	Type   string `json:"type"`
	SERVER string `json:"SERVER"`
	TS     string `json:"TS"`
}

func LBankWsConnect(symbolList []string) {

	if len(symbolList) <= 0 {
		log.Panic(errors.New("Okex订阅的交易对数量为空"))
	}

	ws, err := websocket.Dial(LBankWsUrl, "", LBankOrigin)

	if err != nil {
		log.Println(err.Error())
		return
	}

	//循环订阅交易对
	for _, symbol := range symbolList {
		message := "{\"action\": \"subscribe\", \"subscribe\": \"trade\", \"pair\": \"" + symbol + "\"}"

		_, err = ws.Write([]byte(message))
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Printf("订阅: %s \n", message)
	}

	//统计连续错误次数
	var readErrCount = 0
	var msg = make([]byte, LBankBufferSzie)
	for {
		if readErrCount > LBankErrorLimit {
			ws.Close()
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
		revMsg := string(msg[:m])
		log.Printf("LBank接收：%s \n", revMsg)
		if strings.Contains(revMsg, "ping") {
			var ping map[string]string
			json.Unmarshal(msg[:m], &ping)
			pongStr := "{\"action\": \"pong\", \"pong\": \"" + ping["ping"] + "\"}"
			log.Println("ping消息回应", pongStr)
			ws.Write([]byte(pongStr))
			continue
		}
		var t TradeDetail
		err = json.Unmarshal(msg[:m], &t)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Lbank输出对象", t)
	}
}
