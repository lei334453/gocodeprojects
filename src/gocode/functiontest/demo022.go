package main

import "fmt"

func test (num int){
	fmt.Println(num)
}
// func test03(num1 int,num2 float32,testFunc func(int)){
//   fmt.Println("test-3========")
// }
type myfunc func(int)//函数别名
func test03(num1 int,num2 float32,testFunc myfunc){
  fmt.Println("test-3========")
}
//求两个数的和和差
func test04(num1 int ,num2 int)(int,int){
	result01:=num1+num2
	result02:=num1-num2
     return result01,result02
}
//好的写法 返回值命名 返回值不需要对应
func test04(num1 int ,num2 int)(sum int,sub int){
	sum:=num1+num2
	sub:=num1-num2
     return sum,sub
}
func main(){
	a:= test
	fmt.Println("a类型是:%T,test函数的类型是:%T",a,test)
     a(10)
	test03(10,3.42,test)
	test03(10,4.23,a)
//别名
	type myInt int
	var num11 myInt = 20
    fmt.Println("num1",num11)
  var num22 int = 30
//   num22=num1虽然是别名在go里是不符合的
  num22=int(num11)//可以强转为int这样比较合适
  fmt.Println(num22)

  fmt.Println("========================")
  test03(100,2.32,a)
}