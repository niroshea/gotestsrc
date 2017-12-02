package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func send(ino int, conn net.Conn) {
	for {
		var inStr string
		//读取屏幕输入
		fmt.Scanln(&inStr)
		//var data = "客户端：" + strconv.Itoa(ino) + "。发送信息：" + strconv.Itoa(index)

		if inStr == "/quit" {
			log.Printf("Info: %s\n", "ByeBye")
			conn.Close()
			os.Exit(0)
		}
		//信息发送至服务器
		len, err := conn.Write([]byte(inStr))
		//fmt.Println("客户端接收... ", len, " 字节数据。")
		if err != nil {
			fmt.Println(len, err, "Close Connct")
			conn.Close()
			return
		}
		//time.Sleep(2 * time.Second)
	}

}

func recive(re int, conn net.Conn) {

	for {
		var bufr = make([]byte, 180)
		//var bufr []byte
		//读取服务器发送过来的信息
		len, e := conn.Read(bufr)
		fmt.Println("对方说：" + string(bufr))
		if e != nil {
			fmt.Println(len, e, "Close Connct")
			conn.Close()
			return
		}
		//index = index + 1
		time.Sleep(1 * time.Second)
	}
}

func main() {
	fmt.Println("client main start")

	conn, e := net.Dial("tcp", "127.0.0.1:7878")
	if e != nil {
		panic("e is nil")
	}

	//1个goroutine用于处理信息
	for index := 0; index < 1; index++ {
		go send(index, conn)
		go recive(index+10, conn)
	}
	for {
		time.Sleep(1e9)
	}
}
