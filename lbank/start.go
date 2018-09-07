package lbank

import (
	"ccxt/utils"
	"log"
	"encoding/json"
)

func StartWs(proxy string , useProxy bool) {

	if useProxy && proxy != "" {
		utils.UseProxy = useProxy
		utils.ProxyUrl = proxy
	}

	resp := utils.HttpRequest(LBankSymbole , "GET" , map[string]string{"contentType":"application/x-www-form-urlencoded"})
	var symboleList []string
	err := json.Unmarshal([]byte(resp) , &symboleList)

	if err != nil{
		log.Panic(err)
	}
	log.Println(symboleList)
	go LBankWsConnect(symboleList)
}
