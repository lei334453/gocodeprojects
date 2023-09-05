package main

import "fmt"

func main() {
	// map的创建方式
	//map的操作
	b := make(map[int]string)
	b[2023821]="粘单"
	b[2023831]="张sa"
	b[2023831]="xiaolei"//修改更新
	b[2023433]="哈哈"//增加
	fmt.Println(b)
	//删除
	// delete内置函数
	delete(b,2023433)
	
	fmt.Println(b)

	//查找
	value,flag:=b[2023821]
	fmt.Println(value,flag)
	value2,flag2:=b[32321]
	fmt.Println(value2,flag2)

}