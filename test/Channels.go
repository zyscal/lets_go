package main

import (
	"fmt"
	"time"
)

func main()  {
	channel_x := make(chan int)
	channel_xx := make(chan int)
	done := make(chan bool)
	go create_num(channel_x, done)
	go squarer(channel_x, channel_xx)
	go print_ans(channel_xx)
	flag := <- done
	fmt.Println("flag : ", flag)
}
func create_num(channel_x chan int, done chan bool)  {
	for i := 0; i < 5; i++{
		channel_x <- i
		close(channel_x)
		time.Sleep(2 * time.Second)
	}

	done <- true

}
func squarer(channel_x, channel_xx chan int)  {
	for{
		x, ok := <- channel_x
		if !ok{
			fmt.Println("the channel is closed")
			break;
		}
		channel_xx <- x*x
	}

}
func print_ans(channel_xx chan int)  {
	for{
		xx :=<- channel_xx
		fmt.Println("the ans is ： ", xx)
	}

}
