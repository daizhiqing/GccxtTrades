package gateio

import (
	"ccxt/model"
	"ccxt/utils"
	"strconv"

	"github.com/sirupsen/logrus"
)

//格式化为统一的输出结构
func DataParser(t TradeDetail, id int) {

	var commonData = model.TradeTransData{}

	commonData.Exchange = Name
	commonData.ExchangeId = id
	a, b := utils.FmtSymbol(t.Symbole)

	if a != "" && b != "" {
		commonData.Symbol = a + "/" + b

		for _, tradeTemp := range t.tradeList {
			var trade model.TradeEntity
			trade.Amount = tradeTemp.Amount
			trade.Price = tradeTemp.Price
			trade.Symbol = commonData.Symbol
			trade.Side = tradeTemp.Type
			trade.Timestamp = strconv.Itoa(int(tradeTemp.Time * 1000))

			commonData.Trades = append(commonData.Trades, trade)
		}
		logrus.Infof("输出MQ消息:%s", commonData.ToBody())
		model.DataChannel <- commonData
	}

}
