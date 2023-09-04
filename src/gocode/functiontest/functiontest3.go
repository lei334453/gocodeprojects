package main

import "fmt"

func exchangeNum(num1 int,num2 int)(int,int){
	var t int
	t=num1
	fmt.Println(t)
	num1 = num2
	num2 = t
	return num1,num2
}
//可变参数 支持可变参数
func test(args...int){
	//处理 当成切片处理
	for i:=0;i<len(args);i++{
		fmt.Println(args[i])
	}
  
}

func main(){
	test()
	fmt.Println("--------------")
	test(12,13,132,32)
	fmt.Println("------------")
	//使用函数交互两个值
	var num1 int =10
	var num2 int =20
	fmt.Println("交换前的两个数：num1=%v,num2=%v \n",num1,num2)
	//如果不注意返回值 num1和num2交换后的值只在函数作用域中有用,在main函数中访问她还是访问的是自己的作用域所以没变
	num1,num2=exchangeNum(num1,num2)
	fmt.Println("交换前的两个数：num1=%v,num2=%v \n",num1,num2)
}