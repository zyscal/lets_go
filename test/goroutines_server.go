package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main()  {
	go func() int {
		var i int =5
		fmt.Printf("func 1\n")
		for{}
		return i
	}()
	fmt.Println("after func")
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
//func handleConn(c net.Conn) {
//	defer c.Close()
//	for {
//		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
//		if err != nil {
//			return // e.g., client disconnected
//		}
//		time.Sleep(1 * time.Second)
//	}
//}
func handleConn(c net.Conn)  {
	input := bufio.NewScanner(c)
	for input.Scan(){
		go echo(c, input.Text())
	}
}
func echo(c net.Conn, shout string)  {
	fmt.Fprintln(c, strings.ToUpper(shout))
	time.Sleep(time.Second * 2)
	fmt.Fprintln(c, shout)
	time.Sleep(time.Second * 2)
	fmt.Fprintln(c, strings.ToLower(shout))
}