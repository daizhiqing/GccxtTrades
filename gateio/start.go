package gateio

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

	resp, err := utils.HttpGet(GateioSymbole).Array()
	if err != nil {
		log.Panic(err)
	}
	var symboleList []string
	for _, m := range resp {
		symboleList = append(symboleList, m.(string))
	}
	GateioWsConnect(symboleList)
}
