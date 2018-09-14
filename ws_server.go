package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"golang.org/x/net/websocket"
	"time"
)

func echoHandler(ws *websocket.Conn) {
	msg := make([]byte, 512)
	for {
		if ws.IsClientConn() {
			log.Print("IsClientConn..........")
		}
		if ws.IsServerConn() {
			log.Print("IsServerConn..........")
		}
		n, err := ws.Read(msg)
		if err != nil {
			//log.Fatal(err)
			log.Print(err.Error())
			break
		}
		fmt.Printf("Receive: %s\n", msg[:n])

		send_msg := "[" + string(msg[:n]) + "]"
		_, err = ws.Write([]byte(send_msg))
		if err != nil {
			log.Print(err.Error())
			break
			//log.Fatal(err)
		}
		fmt.Printf("Send: %s\n", send_msg)
	}

}

func mainBak() {
	runtime.GOMAXPROCS(4)
	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main()  {
	//timeStr := "13:53:13"
	////转成时间戳
	//	loc, _ := time.LoadLocation("Asia/Shanghai")
	//nowStr := time.Now().In(loc).Format("2006-01-02 ")
	//tm, err := time.ParseInLocation("2006-01-02 15:04:05", nowStr+timeStr, loc)
	//if err != nil{
	//	panic(err)
	//}
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
}