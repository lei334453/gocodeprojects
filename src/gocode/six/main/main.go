package main

import "fmt"

func main(){
	test()
	fmt.Println("上面的除法执行完毕")
	fmt.Println("正常执行下面的操作。。。。")

}

func test(){
	defer func(){
		//recover 内置函数
		err:=recover()
		//如果没有捕获错误 返回值为零
		if err != nil{
			fmt.Println("错误已经捕获")//error: integer divide by zero
			fmt.Println("err是:",err)//这个捕获错误的代码匿名自调用 expression in defer must be function call
		}
	}()

	num1:=10
	num2:=0
    result:=num1/num2//出错 程序恐慌出错 panic
    fmt.Println(result)
}