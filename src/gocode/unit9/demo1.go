package main

import "fmt"

func main(){
	c :=map[int]string{
		2003322:"裤子",
		2003211:"衣服",
	}
	c[2003343]="卫衣"

	fmt.Println(len(c))

	for k,v :=range c{
		fmt.Printf("key为:%v value为:%v\t",k,v)
	}
	fmt.Println("========================================================")
	//加深难度
	a :=make(map[string]map[int]string)
	//赋值操作
	a["班级一"]=make(map[int]string,4)
	a["班级一"][2004401]="张三"
	a["班级一"][2004402]="丽丽"
	a["班级一"][2004403]="露露"
	a["班级一"][2004404]="嘿嘿"
	a["班级二"]=make(map[int]string,4)
	a["班级二"][20040912]="四十"
	a["班级二"][20040913]="三三"
	a["班级二"][20040914]="萨科"
	a["班级二"][20040915]="小奥"

	for k1,v1 := range a{
		fmt.Println(k1)
		for k2,v2 :=range v1{
			fmt.Printf("学生学号为:%v 学生姓名为:%v\t",k2,v2)
		}
		fmt.Println()
	}
}