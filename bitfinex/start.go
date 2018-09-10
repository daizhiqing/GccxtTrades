package bitfinex

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

	resp, err := utils.HttpGet(BitfinexSymbole).Array()
	if err != nil {
		log.Panic(err)
	}
	var symboleList []string
	for _, m := range resp {
		symboleList = append(symboleList, m.(string))
	}
	log.Println(symboleList)

	go BitfinexWsConnect(symboleList)
}
