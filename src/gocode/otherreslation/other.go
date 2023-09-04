package main

import "fmt"

func main(){
//定义一个变量
//&+变量名 取得地址
//* + 变量 取得值
var age int = 18
fmt.Println("age 对应的存储空间的地址为:",&age)

var ptr *int = &age
// fmt.Println(ptr)
fmt.Println("ptr这个指针指向的具体数值为:",*ptr)

}