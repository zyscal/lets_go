package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main()  {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil{
		fmt.Println("failed to conn")
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
	conn.Close()
}
func mustCopy(dst io.Writer, src io.Reader)  {
	if _, err := io.Copy(dst, src); err != nil{
		fmt.Println(err)
	}
}