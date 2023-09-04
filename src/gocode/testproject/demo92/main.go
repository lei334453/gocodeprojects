package main

import "fmt"

//有局部变量  也有全局变量
var n7 = 234
var n8 = 76

//简洁的方式
var (
	n9  = 500
	n10 = 300
)

func main() {
	// fmt.Println("hello")
	//    第一种使用情况  默认值
	var a int
	fmt.Println("a=", a)
	//第三种
	var num = 20
	fmt.Println("num=", num)
	//第四种 省略
	sex := "男"
	fmt.Println("sex=", sex)

	//  声明多个变量
	var p1, p2, p3 int
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
	fmt.Println("p3=", p3)

	var n4, name, n5 = 10, "jack", 7.8
	fmt.Println(n4)
	fmt.Println(name)
	fmt.Println(n5)

	fmt.Println(n10, n9)
	fmt.Println(n7, n8)
}
