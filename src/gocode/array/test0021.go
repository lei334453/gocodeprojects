package main

import "fmt"

func main(){

	// var arr1 =[3]int{3,4,5}
	// fmt.Println(arr1)
	// fmt.Printf("数组的类型:%T",arr1)
    // var arr2 =[8]int{3,4,5,7,7,8,9}
	// fmt.Println(arr2)
	// fmt.Printf("数组的类型:%T",arr2)
   var arr3=[3] int {3,4,6}
   test1(arr3)
fmt.Println(arr3)//[3,4,6]

}

func test1(arr [3]int){
	arr[0]=7
}