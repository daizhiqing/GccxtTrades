package binance

import (
	"GccxtTrades/utils"

	log "github.com/sirupsen/logrus"
)

func StartWs(proxy string, useProxy bool) {
	if useProxy {
		utils.UseProxy = useProxy
		if proxy != "" {
			utils.ProxyUrl = proxy
		}
	}
	//获取火币的所以交易对
	binanceSymbols := utils.HttpGet(BinanceSymbole).Get("symbols").MustArray()

	var syList []string
	for _, m := range binanceSymbols {
		str := m.(map[string]interface{})["symbol"].(string)
		syList = append(syList, str)
	}
	log.Println(len(syList), syList)
	go BinanceWsConnect(syList)
}
