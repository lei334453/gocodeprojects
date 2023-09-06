package main
import "fmt"
type Student struct{
	Age int
}
type Person struct{
	Age int
}
func main(){
	var s Student = Student{10}
	var p Person = Person{10}
	// s=p//直接赋值不行
	s = Student(p)//转换
	fmt.Println(s)
	fmt.Println(p)
}