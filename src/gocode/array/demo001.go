package main

import "fmt"

func main(){

	var arr1 [3]int =[3]int{3,4,5}
	fmt.Println(arr1)

	//第二种

	var arr2=[3]int{1,5,7}

	fmt.Println(arr2)


	//第三种

	var arr3=[...]int{4,5,6,7}
	fmt.Println(arr3)

	//第四种

	var arr4=[...] int{2:88,0:33,1:44,3:77}
	fmt.Println(arr4)

}