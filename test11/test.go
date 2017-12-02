package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

//创建 动态数组 bufferServer
var buffServer = make([]byte, 1024)

//动态数组 buffClient
var buffClient = make([]byte, 1024)

//map映射，一一映射，key为string类型 值为 net.Conn 类型
var clients = make(map[string]net.Conn)

//缓冲大小为10 的通道
var messages = make(chan string, 10)

//TCPServer 创建监听服务
func TCPServer(port string) {
	port = ":" + port

	//解析端口
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	if err != nil {
		log.Printf("error: %s\n", err)
		return
	}

	//创建监听
	l, err := net.ListenTCP("tcp", tcpAddr)
	defer l.Close()
	if err != nil {
		log.Printf("error: %s\n", err)
		return
	}
	log.Printf("开启了服务端:%s\n", "正在监听客户端的请求")

	//服务器向各个客户端广播消息，如果没有客户端，那就阻塞
	go BoradCast(clients, messages)
	for {
		//死循环

		//返回呼叫监听器的客户端连接
		conn, err := l.Accept()
		if err != nil {
			log.Printf("error: %s\n", err)
		}
		//将刚呼叫我的客户端加入客户端列表
		clients[conn.RemoteAddr().String()] = conn
		go HandleServer(conn, messages)
	}
}

//HandleServer 将连接中的信息写入通道中
func HandleServer(conn net.Conn, messages chan string) {
	for {
		//从连接中读取信息，写入bufferSever
		_, err := conn.Read(buffServer)
		if err != nil {
			log.Printf("error: %s\n", err)
			conn.Close()
			return
		}
		//将bufferSever中的数据转换成字符串 传递给messages通道
		messages <- string(buffServer)
	}
}

//BoradCast 向连接到服务器的客户端广播信息
func BoradCast(clients map[string]net.Conn, messages chan string) {

	for {
		//读取通道中的信息
		msg := <-messages
		for index, client := range clients {
			//服务器向所有客户端连接写入一条msg，该信息由bufferSevers传入 messages 通道
			_, err := client.Write([]byte(msg))
			if err != nil {
				log.Printf("error: %s\n", err)
				//估计是map的一个方法，删除一条键值对
				delete(clients, index)
			}

		}
	}

}

//TCPClient 连接监听服务的客户端
func TCPClient(serverAddr string) {
	//解析信息
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverAddr)
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
	username := conn.LocalAddr().String()
	for {

		//读取屏幕输入
		fmt.Scanln(&input)
		if input == "/quit" {
			log.Printf("Info: %s\n", "ByeBye")
			conn.Close()
			os.Exit(0)
		}
		_, err := conn.Write([]byte(username + ":" + input))
		if err != nil {
			log.Printf("error: %s\n", err)
			return
		}

	}

}

//开启TCPServer: go run main.go server 端口号(8080)
//开启TCPClient: go run main.go client 本地ip地址:端口号(8080)
func main() {
	if len(os.Args) != 3 {
		log.Println("command error")
		os.Exit(0)
	}

	// fmt.Println(tcpAddr.IP)
	// fmt.Println(tcpAddr.Port)
	// fmt.Println(tcpAddr.Zone)

	// os.Exit(0)

	if os.Args[1] == "server" {
		TCPServer(os.Args[2])
	}
	if os.Args[1] == "client" {
		TCPClient(os.Args[2])
	}
}
