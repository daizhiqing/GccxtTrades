package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"golang.org/x/net/websocket"
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
