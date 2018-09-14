package okex

import (
	"ccxt/model"
	"ccxt/utils"
	"strconv"
	"github.com/sirupsen/logrus"
	"strings"
)

//格式化为统一的输出结构
func DataParser(t TradeDetail, id int) {

	var commonData = model.TradeTransData{}

	commonData.Exchange = Name
	commonData.ExchangeId = id

	t.Channel = strings.Replace(t.Channel , "ok_sub_spot_","" , -1)
	t.Channel = strings.Replace(t.Channel , "_deals","" , -1)

	a, b := utils.FmtSymbol(t.Channel)

	if a != "" && b != "" {
		commonData.Symbol = a + "/" + b

		var trade model.TradeEntity

		trade.Amount = t.Amount
		trade.Price = t.Price
		trade.Symbol = commonData.Symbol

		if "ask" == t.Type {
			trade.Side = "sell"
		}
		if "bid" == t.Type {
			trade.Side = "buy"
		}

		trade.Timestamp = strconv.FormatInt(t.Ts , 10)
		commonData.Trades = append(commonData.Trades, trade)

		logrus.Infof("输出MQ消息:%s", commonData.ToBody())
		model.DataChannel <- commonData
	}

}
