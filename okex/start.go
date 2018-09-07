package okex

import (
	"ccxt/utils"
	"ccxt/config"
	"log"
)

func StartWs(proxy string , useProxy bool)  {

	if useProxy {
		utils.UseProxy  = useProxy
		if proxy != ""{
			utils.ProxyUrl = proxy
		}
	}

	okSymbols := utils.HttpGet(config.OkexSymbols).Get("data").MustArray()

	var syList []string

	for _,m := range okSymbols {
		str := m.(map[string]interface{})["symbol"].(string)
		syList = append(syList , str)
	}
	log.Println("okex:" , syList)

	go OkexWsConnect(syList)
}
