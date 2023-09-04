package main

import "fmt"
// func test(num int){
// 	num=30
// 	fmt.Println("test---",num)
// }
// func main(){
// 	var num int = 10
// 	test(num)
// 	fmt.Println("main---",num)
// }
//变成指针
func test(num *int){
	*num=30
	// num=30
	fmt.Println("test---",num)
}
func main(){
	var num int = 10
	fmt.Println(&num)
	// test(num)
	test(&num)//传入num地址
	fmt.Println("main---",num)
}