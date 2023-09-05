package main

import "fmt"

func main(){
	//定义map  只声明不分配空间
	var a map[int]string 
	//make函数进行初始化
	a = make(map[int]string,10)//map可以存放十个键值对
	a[20235432]="张三"
	a[20234564]="李四"
	a[20234423]="王五"
	a[20234421]="王二"
	a[20234421]="王意义"//重复的key
	//map是无序的 初始时需要make key重复 覆盖 value可以重复
	fmt.Println(a)//map[20234423:王五 20234564:李四 20235432:张三]
}