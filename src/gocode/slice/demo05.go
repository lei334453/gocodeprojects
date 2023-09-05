package main
import "fmt"
func main(){
	//定义数组
	var interarr [6]int=[6]int {1,3,5,6,7,8}
	var slice[]int = interarr[1:4]
	fmt.Println(len(slice))

	slice2:=append(slice,88,40)
	fmt.Println(slice2)
	fmt.Println(slice)
	slice=append(slice,88,40)
	fmt.Println(slice)

}