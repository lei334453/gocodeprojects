package main

import "fmt"

func main(){
	var count int = 400
	if count<30 {
		fmt.Println("你的库存不足")
	}else{
		fmt.Println("库存充足")
	}
}