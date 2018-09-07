package utils

import (
	"net/http"
	"io/ioutil"
	"log"
	"net/url"
	"crypto/tls"
	"time"
	"github.com/bitly/go-simplejson"
)
//是否使用代理
const useProxy  = true
//代理服务
const proxyUrl  = "socks5://127.0.0.1:1086"

//设置代理
func proxyReqClient() *http.Client{
	proxy, _ := url.Parse(proxyUrl)
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10, //超时时间
	}
	return client
}

//Get 请求工具类  返回JSON对象
func HttpGet(url string)(*simplejson.Json){

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	if url == "" {
		panic("GET请求url不合法")
	}
	var resp *http.Response
	var err error
	if useProxy{
		resp, err = proxyReqClient().Get(url)
	}else{
		resp, err = http.Get(url)
	}
	CheckErr(err)


	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("请求返回：" , string(body))
	json , err := simplejson.NewJson(body)
	return json
}
