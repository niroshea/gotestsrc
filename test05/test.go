package main

import (
	"bytes"
	"flag"
	//"fmt"
	"io"
	"log"
	"net"
	"time"
)

var arg01 = flag.String("port", "8000", "默认 8000")

//StrApend xxx
func StrApend(buf bytes.Buffer, strs ...string) string {
	for _, str := range strs {
		buf.WriteString(str)
	}
	return buf.String()
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:03:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {

	flag.Parse()
	var buf bytes.Buffer
	buf.WriteString("localhost:")
	buf.WriteString(*arg01)
	rplace := buf.String()

	//string ipport = "localhost:"
	listen, err := net.Listen("tcp", rplace)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}
