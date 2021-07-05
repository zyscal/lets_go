package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main()  {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil{
		fmt.Println("fail to listener")
	}
	for{
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept")
			continue
		}
		handleConn(conn)
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}