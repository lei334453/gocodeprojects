package main

import "fmt"

func main() {
	// map的创建方式
	//方式二
	b := make(map[int]string)
	b[2023821]="粘单"
	b[2023831]="张sa"
	fmt.Println(b)

	//方式三
	c :=map[int]string{
		2003322:"裤子",
		2003211:"衣服",
	}
	c[2003343]="卫衣"
	fmt.Println(c)
}