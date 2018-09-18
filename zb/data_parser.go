package zb

import (
	"GccxtTrades/model"
	"GccxtTrades/utils"
	"strconv"
	"github.com/sirupsen/logrus"
	"strings"
)

//格式化为统一的输出结构
func DataParser(t TradeDetail, id int) {

	var commonData = model.TradeTransData{}

	commonData.Exchange = Name
	commonData.ExchangeId = id

	t.Channel = strings.Replace(t.Channel , "_trades","" , -1)

	a, b := utils.FmtSymbol(t.Channel)

	if a != "" && b != "" {
		commonData.Symbol = a + "/" + b

		for _,td := range t.Data  {
			var trade model.TradeEntity

			trade.Amount = td.Amount
			trade.Price = td.Price
			trade.Symbol = commonData.Symbol
			trade.Side = td.Type
			trade.Timestamp = strconv.Itoa(td.Date*1000 )
			commonData.Trades = append(commonData.Trades, trade)
		}

		logrus.Infof("输出MQ消息:%s", commonData.ToBody())
		model.DataChannel <- commonData
	}

}
