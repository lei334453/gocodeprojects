package main

import "fmt"

func main() {
	//算术运算符
	var n1 int = +10

	fmt.Println(n1)
	var n2 int = 4 + 7
	fmt.Println(n2)
	var str string = "abc" + "def"
	fmt.Println(str)

	//除号
	fmt.Println(10 / 4)   //两个int计算 即使整数
	fmt.Println(10.0 / 3) //浮点数参与运算
	//取模操作%
	fmt.Println(10 % 3)
	fmt.Println(-10 % 3)
	fmt.Println(10 % -3)
	fmt.Println(-10 % -3)
	//++自增 --自减

	var a int = 10
	a++
	fmt.Println(a)
	//在go语言里非常简单 不能参加运算 与js与java不同
	// 不能写在变量之前

}
