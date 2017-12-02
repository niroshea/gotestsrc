package main

import (
	"bytes"
	"log"
	"net"
	"os"
	"strconv"
)

type qsession struct {
	aclient string
	aconn   net.Conn

	bclinet string
	bconn   net.Conn
}

//map映射，一一映射，key为string类型 值为 net.Conn 类型
var clients = make(map[string]net.Conn)

//缓冲大小为10 的通道
var msgChannel = make(chan string, 10)

// session 通道，大小为10
var sessionChannel = make(chan qsession, 10)

func qqServer() {
	//解析端口
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":10901")
	if err != nil {
		log.Printf("解析端口失败: %s\n", err)
		return
	}

	//创建监听
	l, err := net.ListenTCP("tcp", tcpAddr)
	defer l.Close()
	if err != nil {
		log.Printf("创建TCP监听失败: %s\n", err)
		return
	}
	log.Printf("开启了服务端:%s\n", "正在监听客户端的请求")

	go handleSession(clients, sessionChannel)

	count := 0
	for {
		//返回连接到server的Conn
		conn, err := l.Accept()
		if err != nil {
			log.Printf("获取客户端连接失败: %s\n", err)
			continue
		}
		count++
		conname := "conn-" + strconv.Itoa(count)
		//将刚呼叫我的客户端加入客户端列表
		clients[conname] = conn
		//告知 这个客户端他的编号,并且与他进行对话
		go tellNO(conn, conname, clients)
		go createSession(conn, conname, sessionChannel)
		//go hanlechatmsg(conn, msgChannel)
		//if _, ok := clients[conname]; !ok {	}
	}
}

func tellNO(conn net.Conn, conname string, mapx map[string]net.Conn) {
	var buf bytes.Buffer
	//conname = "编号：" + conname + ","
	buf.WriteString("编号: ")
	buf.WriteString(conname)
	buf.WriteString(", 在线的客户有：")

	for index := range mapx {
		buf.WriteString(index)
		buf.WriteString("、")
	}

	if _, err := conn.Write([]byte(buf.String())); err != nil {
		log.Printf("告知客户端其对话编号错误: %s\n", err)
		delete(mapx, conname)
	}
}

func createSession(conn net.Conn, conname string, sessionChannel chan<- qsession) {
	for {
		var buffServer []byte
		//从连接中读取信息，写入bufferSever
		_, err := conn.Read(buffServer)

		log.Println("客户端 " + conname + " 发过来的信息：" + string(buffServer))

		continue
		os.Exit(0)

		if err != nil {
			log.Printf("从连接中读取session NO 出错，连接关闭: %s\n", err)
			conn.Close()
			return
		}
		sessioni := qsession{aclient: conname, aconn: conn, bclinet: string(buffServer)}
		sessionChannel <- sessioni
	}
}

func hanlechatmsg(conn net.Conn, messages chan<- string) {
	for {
		var buffServer []byte
		//从连接中读取信息，写入bufferSever
		_, err := conn.Read(buffServer)
		if err != nil {
			log.Printf("从连接中读取数据出错，连接关闭: %s\n", err)
			conn.Close()
			return
		}
		//将bufferSever中的数据转换成字符串 传递给messages通道
		messages <- string(buffServer)
	}
}

func chatwithfriend(clients map[string]net.Conn, messages <-chan string) {
	for {
		//读取通道中的信息
		msg := <-messages
		//从msg中读取到A客户端发出的信息，分离出这个信息要发送给B客户端的编号

		//查找这个客户段编号，如果没有，则返回没有此用户信息给A客户端，否则发送给B

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

func handleSession(clients map[string]net.Conn, sessionChannel <-chan qsession) {
	for {
		sessioni := <-sessionChannel

		ibconn, ok := clients[sessioni.bclinet]
		if ok {
			sessioni.bconn = ibconn
		} else {
			log.Printf("会话处理失败！\n")
			continue
		}

		//A发送给B
		go func(se qsession) {
			var buf []byte
			_, err1 := se.aconn.Read(buf)
			if err1 != nil {
				log.Printf("Afasonggei B error1: %s\n", err1)
			}
			_, err2 := se.bconn.Write(buf)
			if err2 != nil {
				log.Printf("Afasonggei B error2: %s\n", err2)
			}
		}(sessioni)

		//B发送给A
		go func(se qsession) {
			var buf []byte
			_, err1 := se.bconn.Read(buf)
			if err1 != nil {
				log.Printf("Bfasonggei A error1: %s\n", err1)
			}
			_, err2 := se.aconn.Write(buf)
			if err2 != nil {
				log.Printf("Bfasonggei A error2: %s\n", err2)
			}
		}(sessioni)

	}
}

func main() {
	qqServer()
	//建立监听

	//列出有哪些客户在线

	//你想与谁通话，你通话的客户已下线

	//开始通话

}
