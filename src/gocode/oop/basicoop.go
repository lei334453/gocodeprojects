package main

import "fmt"

type Teacher struct{
	Name string
	Age int
	School string 
}

func main(){
	//处理老师实例
	var t1 Teacher
	fmt.Println(t1)
    t1.Name="xiaolei"
	t1.Age=30
	t1.School="五道口大学"
	fmt.Println(t1)
}