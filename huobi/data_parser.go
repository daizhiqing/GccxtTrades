package huobi

import (
	"ccxt/model"
	"ccxt/utils"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

//格式化为统一的输出结构
func DataParser(t TradeDetail, id int) {

	var commonData = model.TradeTransData{}

	commonData.Exchange = Name
	commonData.ExchangeId = id
	a, b := utils.FmtSymbol(strings.Split(t.Ch, ".")[1])

	if a != "" && b != "" {
		commonData.Symbol = a + "/" + b

		for _, tradeTemp := range t.Tick.Data {
			var trade model.TradeEntity
			trade.Amount = strconv.FormatFloat(tradeTemp.Amount, 'f', -1, 64)
			trade.Price = strconv.FormatFloat(tradeTemp.Price, 'f', -1, 64)
			trade.Symbol = commonData.Symbol
			trade.Side = tradeTemp.Direction
			trade.Timestamp = strconv.Itoa(tradeTemp.Ts)

			commonData.Trades = append(commonData.Trades, trade)
		}
		logrus.Infof("输出MQ消息:%s", commonData.ToBody())
		model.DataChannel <- commonData
	}

}
