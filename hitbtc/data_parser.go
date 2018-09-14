package hitbtc

import (
	"ccxt/model"
	"ccxt/utils"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

//格式化为统一的输出结构
func DataParser(t TradeDetail, id int) {

	var commonData = model.TradeTransData{}

	commonData.Exchange = Name
	commonData.ExchangeId = id
	a, b := utils.FmtSymbol(t.Params.Symbol)

	if a != "" && b != "" {
		commonData.Symbol = a + "/" + b

		for _, tradeTemp := range t.Params.Data {
			var trade model.TradeEntity
			trade.Amount = tradeTemp.Quantity
			trade.Price = tradeTemp.Price
			trade.Symbol = commonData.Symbol
			trade.Side = tradeTemp.Side

			//将英国伦敦时间
			toBeCharge := strings.Replace(tradeTemp.Timestamp, "Z", "", -1)
			timeLayout := "2006-01-02T15:04:05"
			loc, err := time.LoadLocation("GMT")
			if err != nil {
				logrus.Error(err)
				return
			}
			theTime, err := time.ParseInLocation(timeLayout, toBeCharge, loc)
			if err != nil {
				logrus.Error(err)
				return
			}

			trade.Timestamp = strconv.FormatInt(theTime.Unix()*1000, 10)

			commonData.Trades = append(commonData.Trades, trade)
		}
		logrus.Infof("输出MQ消息:%s", commonData.ToBody())
		model.DataChannel <- commonData
	}

}
