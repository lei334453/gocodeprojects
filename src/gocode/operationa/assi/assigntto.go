package main

import "fmt"

func main() {
	var num1 int = 10
	fmt.Println(num1)
	var num2 int = (10+30)%3 - 4
	fmt.Println(num2)

	var num3 int = 10
	num3 += 20
	fmt.Println(num3)

	//交换
	var a int = 8
	var b int = 5
	fmt.Println("a=%v", a, b)

	var f int //声明默认不赋值
	f = a
	a = b
	b = f
	fmt.Println(a, b)
}
