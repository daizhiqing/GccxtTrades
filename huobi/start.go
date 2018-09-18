package huobi

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
	huobiSymbols := utils.HttpGet(HuoBiSymbols).Get("data").MustArray()
	var syList []string
	for _, m := range huobiSymbols {
		// str := m.(map[string]interface{})["symbol"].(string)
		syList = append(syList, m.(map[string]interface{})["symbol"].(string))
	}
	log.Println("huobi:", syList)

	go HuobiWsConnect(syList)
}
