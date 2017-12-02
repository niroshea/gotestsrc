package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println("xxxxxxx")
		log.Fatal(err)
	}
}

var arg01 = flag.String("port", "8000", "默认 8000")

func main() {

	flag.Parse()
	var buf bytes.Buffer
	buf.WriteString("localhost:")
	buf.WriteString(*arg01)
	rplace := buf.String()

	fmt.Println("===========x=x=x=x=")
	conn, err := net.Dial("tcp", rplace)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//os.Stdout  直接关联设备？
	mustCopy(os.Stdout, conn)
}
