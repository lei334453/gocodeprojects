package main
import "fmt"
func main(){
	// var slice[] int
	// fmt.Println(slice)

	var interarr[6]int = [6]int{1,4,5,6,7,8}
   var slice []int  = interarr[1:4]
   fmt.Println(slice[0])
   fmt.Println(slice[1])
   fmt.Println(slice[2])
//    fmt.Println(slice[3])//越界
//切片可以继续切片

slice2:=slice[1:2]
fmt.Println(slice2)

slice2[0]=66
fmt.Println(interarr)
fmt.Println(slice)

}
