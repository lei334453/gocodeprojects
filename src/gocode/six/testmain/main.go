package main

import (
	"fmt"
	"errors"
)

func main(){
	err:=test()
	if err !=nil{
		fmt.Println("自定义错误",err)
		panic(err)
	}
	fmt.Println("上面的除法执行完毕")
	fmt.Println("正常执行下面的操作。。。。")

}

func test()(err error){
	
    
	num1:=10
	num2:=0
	if num2 == 0{
		//抛出自定义错误
		return errors.New("除数不能为0哦~~~")
		
	}else{
		result:=num1/num2//出错 程序恐慌出错 panic
		fmt.Println(result)
		//return 

		return nil
	}
   
}