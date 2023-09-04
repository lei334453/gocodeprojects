package main

import "fmt"

func main(){
	var score int = 89
	if score >=90{
		fmt.Println("你的成绩很优秀")
	}else if score >=80 {//else if score <90 && score>80
        fmt.Println("你的成绩优秀")
	}else if score>=60{
		fmt.Println("你的成绩还好")
	}else{
		fmt.Println("你的成绩未及格,同学你任然需要努力")
	}
}