package gateio

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

	resp, err := utils.HttpGet(GateioSymbole).Array()
	if err != nil {
		log.Panic(err)
	}
	var symboleList []string
	for _, m := range resp {
		symboleList = append(symboleList, m.(string))
	}
	log.Info("这是一个INFO")
	log.Debug("这是一个DEBUG")

	log.Warn("这是一个WARN")
	log.Error("这是一个ERROR")

}
