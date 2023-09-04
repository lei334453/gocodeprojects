package main

import "fmt"

func main(){
	score1:=34
	score2:=92
	score3:=64
	score4:=79
	score5:=60

	sum:=score1+score2+score3+score4+score5

	avg:=sum/5

	fmt.Println("成绩总和为:%v,成绩的平均数为:%v",sum,avg)
}