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
	mustCopy(os.Stdout, conn)

}
func mustCopy(dst io.Writer, src io.Reader)  {
	fmt.Println("enter into mustcopy")
	if _, err := io.Copy(dst, src); err != nil{
		fmt.Println("enter into if")
		fmt.Println(err)
	}
	fmt.Println("before out of mustcopy")
}