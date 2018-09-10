package huobi

import (
	"ccxt/utils"

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
		str := m.(map[string]interface{})["base-currency"].(string) + m.(map[string]interface{})["quote-currency"].(string)
		syList = append(syList, str)
	}
	log.Println("huobi:", syList)

	go HuobiWsConnect(syList)
}
