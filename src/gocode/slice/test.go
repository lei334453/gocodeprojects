package main
import "fmt"
func main(){
	// var inarr [6]int=[6]int {1,2,3,4,5,7}
   //第二种方式
	slice :=make([]int,4,67)
	fmt.Println("切片的长度:",len(slice))
	fmt.Println("切片的容量:",cap(slice))

	slice[0]=55
    slice[1]=33
    fmt.Println(slice)


	//i第三种方式
	slice2:=[]int{1,4,5}
	fmt.Println(slice2)
	fmt.Println("切片的长度:",len(slice2))
	fmt.Println("切片的容量:",cap(slice2))
	//切片在数组之上
	//定义切片 slice的动态数组
	//切出的一个片段(不包含)
	// var slice[]int = inarr[1:3]
	//简写
	// slice:= inarr[1:3]
	

// 	fmt.Println("inarr",inarr)

// 	fmt.Println("slice:",slice)

// 	fmt.Println("slice的元素个数",len(slice))
// 	fmt.Println("slice的元素个数",cap(slice))


// 	fmt.Printf("数组中下标为1的位置:%p",&inarr[1])
// 	fmt.Printf("slicei下标1的位置,%p ",&slice[0])
// 	slice[1]=18
// 	fmt.Println("inarr:",inarr)
// 	fmt.Println("slice:",slice)
// 	fmt.Printf("slicei下标2的位置,%p ",&slice[1])

    


}