package main

import "fmt"

func sumtwo(a int,b int)(int){
	var sum int=0
	sum+=a
	sum+=b
	return sum	
}
func main(){
	// var num1 int=20
	// var num2 int=30
	// var sum int=0
	//函数的定义样式
	//前括号形参 
	
	sum:=sumtwo(100,200)
	fmt.Println(sum)

	sum1:=sumtwo(100,400)
	fmt.Println(sum1)
}

