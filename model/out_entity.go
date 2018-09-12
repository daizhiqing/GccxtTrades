package model

import (
	"ccxt/config"
	"encoding/json"

	"github.com/sirupsen/logrus"
)

var DataChannel = make(chan TradeTransData, config.ChannelSize)

type TradeEntity struct {
	Symbol string `json:"symbol"`
	//   datetime: 2018-09-12T09:47:05.000Z,
	Side      string `json:"side"`
	Amount    string `json:"amount"`
	Price     string `json:"price"`
	Timestamp string `json:"timestamp"`
}

type TradeTransData struct {
	ExchangeId int           `json:"exchangeId"`
	Symbol     string        `json:"symbol"`
	Trades     []TradeEntity `json:"trades"`
	Exchange   string        `json:"exchange"`
}

func (t *TradeTransData) ToBody() []byte {
	body, err := json.Marshal(t)
	if err != nil {
		logrus.Errorln(err)
		return nil
	}
	return body
}
