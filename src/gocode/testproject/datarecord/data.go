package main

import "fmt"

func main() {
	// int8 1个字节  -2^7~-2^7-1
	// int16
	//int32
	//int64
	// 01111111---->2进制==>转为10进制 127
	//转过来 127---128

	// fmt.Println(127)
	var num1 int8 = 23 //230//溢出
	fmt.Println(num1)

}
