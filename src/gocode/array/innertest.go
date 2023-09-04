package main

import "fmt"

func main(){
	//只是声明
	var arr[3] int16
	//数组长度
	fmt.Println(len(arr))
	fmt.Println(arr)//[0,0,0]
	fmt.Println("arr的地址为:%p",&arr)
	fmt.Println("arr0的地址为:%p",&arr[0])
	fmt.Println("arr1的地址为:%p",&arr[1])
	fmt.Println("arr2的地址为:%p",&arr[2])
	//arr的地址为:%p &[0 0 0]
	//0、1、2、3、4、5、6、7、8、9和字母A、B、C、D、E、F（a、b、c、d、e、f）
	// arr0的地址为:%p 0xc00000a0c8 8
	// arr1的地址为:%p 0xc00000a0ca 8+2
	// arr2的地址为:%p 0xc00000a0cc 8+4
	arr[0]=10
	arr[1]=20
	arr[2]=30

}