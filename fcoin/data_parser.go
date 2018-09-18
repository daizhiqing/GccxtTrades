package fcoin

import (
	"GccxtTrades/model"
	"GccxtTrades/utils"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

//格式化为统一的输出结构
func DataParser(t TradeDetail, id int) {

	var commonData = model.TradeTransData{}

	commonData.Exchange = Name
	commonData.ExchangeId = id

	var trades []model.TradeEntity
	var trade model.TradeEntity
	trade.Amount = strconv.FormatFloat(t.Amount, 'f', -1, 64)
	trade.Price = strconv.FormatFloat(t.Price, 'f', -1, 64)
	if t.Amount >= 0 {
		trade.Side = "buy"
	} else {
		trade.Side = "sell"
	}
	if len(strings.Split(t.Type, ".")) <= 1 {
		return
	}
	a, b := utils.FmtSymbol(strings.Split(t.Type, ".")[1])
	if a != "" && b != "" {
		trade.Symbol = a + "/" + b
		trade.Timestamp = strconv.Itoa(t.Ts)
		trades = append(trades, trade)

		commonData.Symbol = a + "/" + b
		commonData.Trades = trades

		logrus.Infof("输出MQ消息:%s", commonData.ToBody())
		model.DataChannel <- commonData
	}
}
