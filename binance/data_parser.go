package binance

import (
	"ccxt/model"
	"ccxt/utils"
	"strconv"

	"github.com/sirupsen/logrus"
)

//格式化为统一的输出结构
func DataParser(t TradeDatail, id int) {

	var commonData = model.TradeTransData{}

	commonData.Exchange = Name
	commonData.ExchangeId = id

	var trades []model.TradeEntity
	var trade model.TradeEntity
	trade.Amount = t.Data.Quantity
	trade.Price = t.Data.Price
	if t.Data.Buy {
		trade.Side = "buy"
	} else {
		trade.Side = "sell"
	}
	a, b := utils.FmtSymbol(t.Data.Symbol)
	if a != "" && b != "" {
		trade.Symbol = a + "/" + b
		trade.Timestamp = strconv.Itoa(t.Data.Ts)
		trades = append(trades, trade)

		commonData.Symbol = a + "/" + b
		commonData.Trades = trades
	}
	logrus.Infof("输出MQ消息:%s", commonData.ToBody())
	model.DataChannel <- commonData
}
