package lbank

import (
	"GccxtTrades/utils"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func StartWs(proxy string, useProxy bool) {

	if useProxy {
		utils.UseProxy = useProxy
		if proxy != "" {
			utils.ProxyUrl = proxy
		}
	}

	resp := utils.HttpRequest(LBankSymbole, "GET", map[string]string{"contentType": "application/x-www-form-urlencoded"})
	var symboleList []string
	err := json.Unmarshal([]byte(resp), &symboleList)

	if err != nil {
		log.Panic(err)
	}
	log.Println(symboleList)
	go LBankWsConnect(symboleList)
}
