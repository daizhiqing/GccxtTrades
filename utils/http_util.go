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
var UseProxy = false
//代理服务
var ProxyUrl = "socks5://127.0.0.1:1086"

//设置代理
func proxyReqClient() *http.Client{
	proxy, _ := url.Parse(ProxyUrl)
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
	if UseProxy {
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

// method：   GET || POST   headers:自定义的头部
func HttpRequest(webUrl string , method string ,headers map[string]string) string{
	/*
		1. 代理请求
		2. 跳过https不安全验证
		3. 自定义请求头 User-Agent
	*/
	request, _ := http.NewRequest(method, webUrl, nil)

	for key,value := range headers{
		request.Header.Set(key ,value)
	}
	//request.Header.Set("Connection", "keep-alive")
	//request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	proxy, _ := url.Parse(ProxyUrl)
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10, //超时时间
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Println("请求出错了", err)
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return (string(body))
}