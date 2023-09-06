package main
//方法是作用在指定类型数据类型上，和指定的数据类型绑定 因此自定义类型 都可以是方法 而不仅仅是struct
//其实是值传递
import "fmt"
type Person struct {
	
	Name string
}
//名字随意起
func (p Person)test(){
	p.Name="露露"
	fmt.Println(p.Name)
}
func main(){
 var p Person
 p.Name="lili"
 p.test()
 fmt.Println(p.Name)
}