package main

import "fmt"

func main(){
	//逻辑与 只要有一侧是false即使false 
	fmt.Println(true&&true)
	fmt.Println(true&&false)
	fmt.Println(false&&true)
	fmt.Println(false&&false)
    //或逻辑
	fmt.Println(true||true)
	fmt.Println(true||false)
	fmt.Println(false||true)
	fmt.Println(false||false)


	fmt.Println(!true)
	fmt.Println(!false)

}