package main

import "fmt"

//浮点数底层储存  符号位+指数位+尾数位  可能有精度损失

func main() {
	//
	var num1 float32 = 3.14
	fmt.Println(num1)
	var num2 float32 = -3.14
	fmt.Println(num2)
	var num3 float32 = 3.14e-2
	fmt.Println(num3)
	var num4 float32 = 3.14e+2
	fmt.Println(num4)
	//可以表示 正负浮点数
	//浮点数可以用十进制表示形式 也可以科学计数法
	var num5 float64 = 3.14e+2
	fmt.Println(num5)
	var num6 float64 = 3.14e+2
	fmt.Println(num6)

	var num7 float32 = 246.00003223
	var num9 float64 = 32.000000000000019
	fmt.Println(num7)
	fmt.Println(num9)

	var num10 = 3.13
	fmt.Printf("默认的数据类型:%T", num10)

}
