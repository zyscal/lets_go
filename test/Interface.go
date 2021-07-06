package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width float64
	height float64
}
func (r Rect) Area() float64{
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {//必须要实现全部接口
	return 2 * (r.height + r.width)
}

func main()  {
	var s Shape = Rect{1,2}
	fmt.Println("value of s is : ", s)
	fmt.Printf("the type of s is : %T", s)
	fmt.Println("the area() of s is : ",s.Area())
	//fmt.Println("the perimeter() of s is ： ", s.Perimeter())
	
}