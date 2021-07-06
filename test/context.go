package main

import (
	"fmt"
	"time"
)

func main()  {
	go test_more_goroutines(1)
	for{}
}
func test_more_goroutines(id int)  {
	i := 0
	for{
		fmt.Println("id :", id, "\ti : ", i)
		i++
		time.Sleep(time.Second * 2)
		if i == 3{
			go test_more_goroutines(id + 1)
		}
		if i== 7{
			fmt.Println(id, "finish")
			break
		}
	}

}
