package main
import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"bufio"
	"time"
)

func main() {
	var vals []int = []int{1,2,3}
	fmt.Println(sum(vals ...))

	//for {
	//	for _, r := range `-\|/` {
	//		fmt.Printf("\rc", r)
	//		time.Sleep(100 * time.Millisecond)
	//	}
	//}
	//test_spinner()
	//test_http_get()

}
func test_spinner(){
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}


func test_http_get()  {
	uri := "http://www.baidu.com"
	resp, err := http.Get(uri)
	if err != nil{
		fmt.Println("there is something wrong")
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s", b)
}

func test_count_line()  {
	//var s, sep string
	for i := 0; i < len(os.Args); i++{
		fmt.Println(os.Args[i])
	}
	//s := strings.Join(os.Args[1:], "11")
	counts := make(map[string]int)
	file_name := os.Args[1]

	f, err := os.Open(file_name)
	if err != nil {
		fmt.Println("the file name is : ", file_name)

	}

	inputs := bufio.NewScanner(f)
	for inputs.Scan() {
		fmt.Println("--------")
		fmt.Println(inputs.Text())
		counts[inputs.Text()]++
	}
	fmt.Println(counts)

}
func sum(val ...int) int  {
	sum := 0
	for _, tem_val := range val{
		sum += tem_val
	}
	return sum

}