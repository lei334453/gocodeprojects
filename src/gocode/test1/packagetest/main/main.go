package main

// import "fmt"

// //注意细节
// import "gocode/test1/packagetest/utils"
//注意同名与不同名字的情况
import(
	"fmt"
	"gocode/test1/packagetest/utils"
	"gocode/test1/packagetest/cal"
)

func main(){
		fmt.Println("hello")
		a:=aaa.Cal(20,30)
		utils.AddNum()
		fmt.Println(a)
}