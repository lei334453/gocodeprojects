package main
import "fmt"
func main(){
	slice :=make([]int,4,18)
	slice[0]=22
	slice[1]=23
	slice[2]=25
	slice[3]=27
	//普通for循环
	for i:=0;i<len(slice);i++{
		fmt.Printf("slice[%v]=%v\t",i,slice[i])
	}
	fmt.Println("========================================")

	//方式二
	for i,v:=range slice {
		fmt.Printf("下标:%v,元素:%v\n",i,v)
	}
}
