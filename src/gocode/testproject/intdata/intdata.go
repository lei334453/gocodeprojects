package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//无符号整数类型
	/***
	int 有符号
	unit 无字符
	rune 无字符


	*/

	//uint16 2^7+127=255
	var num2 uint8 = 28
	fmt.Println(num2)

	var num3 = 28
	fmt.Printf("num3的类型是: %T", num3)
	fmt.Println()
	fmt.Println(unsafe.Sizeof(num3)) //int8
	//表示学生年纪

	var age byte = 18
	fmt.Println(age)

}
