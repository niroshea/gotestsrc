package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

var connectMap = make(map[int]net.Conn)

type sessionx struct {
	A net.Conn
	B net.Conn
}

func byteToStr(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

//为监听进程 获取连接
func handleAccept(L net.Listener) {
	count := 1024
	//循环accept，循环获取connection
	for {
		//fmt.Println("Accept index=", index)
		conn, err := L.Accept()
		if err != nil {
			fmt.Println("Accept index quit.")
			conn.Close()
			break
		}
		connectMap[count] = conn
		fmt.Println("客户编号：", count)
		if len(connectMap) > 1 {
			var s sessionx
			go createSession(conn, count, connectMap, &s)
		}
		count++
		//go handleConnect(conn, count)

	}
}

func createSession(Ac net.Conn, index int, cm map[int]net.Conn, s *sessionx) {
	fmt.Println("进入创建session")
	atob := make(chan string, 10)
	btoa := make(chan string, 10)
	for {
		var bufr = make([]byte, 30)

		len, e := Ac.Read(bufr)
		mstr := byteToStr(bufr)
		fmt.Println("创建session，想与 " + mstr + "聊天。")
		x, err := strconv.Atoi(mstr)
		fmt.Println("======================================>", x, "   len ", len, " mstr", mstr)
		if e != nil {
			fmt.Println(len, e, "Close Connct")
			Ac.Close()
			delete(cm, index)
			return
		}
		Bc, ok := cm[x]
		if err == nil && ok == true {
			s.A = Ac
			s.B = Bc
			break
		} else {
			fmt.Println("请重新输入...")
		}
	}

	//两个客户端开始对话 a-b
	go funcAtob(s, btoa, atob)
	go funcBtoa(s, btoa, atob)
}

func funcAtob(s *sessionx, btoa chan string, atob chan string) {
	go reciveMsg(s.A, 1, atob)
	go sendMsg(s.B, 2, atob)
}

func funcBtoa(s *sessionx, btoa chan string, atob chan string) {
	go reciveMsg(s.B, 2, btoa)
	go sendMsg(s.A, 1, btoa)
}

func handleConnect(c net.Conn, ha int) {
	//go reciveMsg(c, ha)
	//go sendMsg(c, ha)
}

//服务器发送信息给客户端
func sendMsg(conn net.Conn, se int, msg <-chan string) {
	//var index = 0
	for {
		//var data = "服务器：" + strconv.Itoa(se) + "。发送信息：当前在线人数 " + strconv.Itoa(len(connectMap))
		data := <-msg
		//信息发送至客户端
		len, e := conn.Write([]byte(data))
		fmt.Println("发送 " + strconv.Itoa(se) + " 客户端：" + data)
		//fmt.Println("服务器段发送了... ", len, " 字节数据")
		if e != nil {
			fmt.Println(len, e, "Close Connct")
			conn.Close()
			return
		}
		//index = index + 1
		//time.Sleep(1 * time.Second)
	}
}

//服务器从客户端接收信息
func reciveMsg(conn net.Conn, re int, msg chan<- string) {
	for {
		var bufr = make([]byte, 180)

		len, e := conn.Read(bufr)
		mstr := string(bufr)
		fmt.Println("接收 " + strconv.Itoa(re) + " 客户端：" + mstr)

		if e != nil {
			fmt.Println(len, e, "Close Connct")
			conn.Close()
			return
		}
		msg <- mstr
		time.Sleep(1 * time.Second)
	}
}

func ca() {
	a := recover()
	fmt.Println("Catch", a)
}

func main() {
	defer ca()

	fmt.Println("main start")

	//创建一个listener，用于监听
	L, e := net.Listen("tcp", "127.0.0.1:7878")
	if e != nil {
		fmt.Println("this is null")
		panic("claa null")
	}

	//创建10个goroutine用于accept，在高并发的情况下，accept越多越好
	//for index := 0; index < 2; index++ {

	go handleAccept(L)
	//}

	for {
		time.Sleep(1e9)
	}

}
