package main

import "fmt"

func main(){
var score [5]int //这里是数组定义  var 数组名字[数组大小] 数组类型
    score[0]=34
	score[1]=92
	score[2]=64
	score[3]=79
	score[4]=60

	// sum:=score1+score2+score3+score4+score5
	//求和
	sum:=0
	for i:=0;i<len(score);i++{
       sum+=score[i]
	}

	avg:=sum/5

	fmt.Println("成绩总和为:%v,成绩的平均数为:%v",sum,avg)

}