package main

import "fmt"
//优化代码 键盘录入
func main(){
var score [5]int //这里是数组定义  var 数组名字[数组大小] 数组类型
    // score[0]=34
	// score[1]=92
	// score[2]=64
	// score[3]=79
	// score[4]=60
	for i:=0;i<len(score);i++{
		fmt.Printf("请录入第%d个学生的成绩数据",i+1)
		fmt.Scanln(&score[i])//录入数据  终端输入
	}

	// sum:=score1+score2+score3+score4+score5
	//求和
	// sum:=0
	// for i:=0;i<len(score);i++{
    //    sum+=score[i]
	// }

	// avg:=sum/5

	// 展示每个学生成绩
	//方式一
	// for i:=0;i<len(score);i++{
    // //    sum+=score[i]
	//    fmt.Printf("第%d个学生的成绩为:%d\n",i+1,score[i])
	// }
	//方式二 for :=range 遍历数据  key,value名字随意取
	for key,value :=range score{
		fmt.Printf("第%d个学生的成绩为:%d\n",key+1,value)
	}

	// fmt.Println("成绩总和为:%v,成绩的平均数为:%v",sum,avg)

}