package main

import "fmt"

func main(){
	// var i1 int =1
	// var i2 int =2
	// var i3 int =3
	// var i4 int =4
	// var i5 int =5

	// var sum int=0
	// sum=i1+i2+i3+i4+i5
	// fmt.Println("sum=",sum)
//利用for循环
var sum int=0
for  i:=1; i<=5;i++{
	sum+=i
}
fmt.Println(sum)


}