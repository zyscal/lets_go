package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("count down")
	tick := time.Tick(1 * time.Second)
	for{
		test_tick :=<- tick
		fmt.Println(test_tick)
	}

}
