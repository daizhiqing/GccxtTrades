package lbank

import (
	"ccxt/model"
	"ccxt/utils"
	"strconv"
	"github.com/sirupsen/logrus"
	"time"
)

//格式化为统一的输出结构
func DataParser(t TradeDetail, id int) {

	var commonData = model.TradeTransData{}

	commonData.Exchange = Name
	commonData.ExchangeId = id
	a, b := utils.FmtSymbol(t.Pair)

	if a != "" && b != "" {
		commonData.Symbol = a + "/" + b

			var trade model.TradeEntity
			trade.Amount = strconv.FormatFloat(t.Trade.Amount, 'f', -1, 64)
			trade.Price = strconv.FormatFloat(t.Trade.Price, 'f', -1, 64)
			trade.Symbol = commonData.Symbol
			trade.Side = t.Trade.Direction


		timeLayout := "2006-01-02T15:04:05"
		loc, err := time.LoadLocation("BJT")
		if err != nil {
			logrus.Error(err)
			return
		}
		theTime, err := time.ParseInLocation(timeLayout, t.Trade.TS, loc)
		if err != nil {
			logrus.Error(err)
			return
		}
		trade.Timestamp = strconv.FormatInt(theTime.Unix(), 10)

		commonData.Trades = append(commonData.Trades, trade)
		logrus.Infof("输出MQ消息:%s", commonData.ToBody())
		model.DataChannel <- commonData
	}

}
