package config

import (
	"encoding/json"
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/sirupsen/logrus"
)

const (
	//最大读取错误限制次数
	DEFAULT = 100

	DEFAULT_BUFFER_SIZE = 1024

	ChannelSize = 1024

	MqExchange = "ex-api-mq"

	QueuePre = "trades_"

	ExchangeMap = "[{\"exchangeName\":\"binance\",\"id\":11},{\"exchangeName\":\"huobipro\",\"id\":12},{\"exchangeName\":\"\",\"id\":13},{\"exchangeName\":\"bitfinex\",\"id\":14},{\"exchangeName\":\"\",\"id\":15},{\"exchangeName\":\"okex\",\"id\":16},{\"exchangeName\":\"\",\"id\":17},{\"exchangeName\":\"\",\"id\":18},{\"exchangeName\":\"hitbtc2\",\"id\":19},{\"exchangeName\":\"bithumb\",\"id\":20},{\"exchangeName\":\"\",\"id\":21},{\"exchangeName\":\"\",\"id\":22},{\"exchangeName\":\"zb\",\"id\":23},{\"exchangeName\":\"\",\"id\":24},{\"exchangeName\":\"\",\"id\":25},{\"exchangeName\":\"gateio\",\"id\":60},{\"exchangeName\":\"kucoin\",\"id\":61},{\"exchangeName\":\"bittrex\",\"id\":62},{\"exchangeName\":\"coinsuper\",\"id\":63},{\"exchangeName\":\"fcoin\",\"id\":64},{\"exchangeName\":\"hadax\",\"id\":65},{\"exchangeName\":\"lbank\",\"id\":66}]"
)

func GetExchangeId(exName string) int {
	jsonOb, err := simplejson.NewJson([]byte(ExchangeMap))
	if err != nil {
		logrus.Errorln(err)
		return 0
	}
	arr, err := jsonOb.Array()
	if err != nil {
		logrus.Errorln(err)
		return 0
	}

	for _, m := range arr {
		if m.(map[string]interface{})["exchangeName"] == exName {
			a, _ := strconv.Atoi(m.(map[string]interface{})["id"].(json.Number).String())
			return a
		}
	}
	return 0
}
