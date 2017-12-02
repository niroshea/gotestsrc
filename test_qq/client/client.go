package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

//动态数组 buffClient
var buffClient = make([]byte, 256)

func qqClient() {
	//解析信息
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:10901")
	if err != nil {
		log.Printf("error: %s\n", err)
		return
	}

	//与服务器进行通信建立连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	if err != nil {
		log.Printf("error: %s\n", err)
		return
	}
	log.Printf("客户端:%s开始连接\n", conn.LocalAddr().String())
	go ClientSend(conn)
	for {
		//读取连接中接收到的信息
		_, err := conn.Read(buffClient)
		if err != nil {
			log.Printf("error: %s\n", err)
			return
		}
		log.Printf("收到消息:%s\n", string(buffClient))
	}
}

//ClientSend 客户端向服务器发送信息
func ClientSend(conn net.Conn) {

	var input string
	//username := conn.LocalAddr().String()
	for {
		//读取屏幕输入
		fmt.Scanln(&input)
		if input == "/quit" {
			log.Printf("Info: %s\n", "ByeBye")
			conn.Close()
			os.Exit(0)
		}
		_, err := conn.Write([]byte(input))
		if err != nil {
			log.Printf("error: %s\n", err)
			return
		}

	}

}

func main() {
	qqClient()
}
