package main

import "fmt"

type integer int

func (i integer) print(){
	fmt.Println("i=",i)
}

func (i *integer) change(){
	*i=30
	fmt.Println("i=",*i)
}

func main(){
	var i integer =20
	i.print()
	i.change()
	fmt.Println("i=",i)
}
