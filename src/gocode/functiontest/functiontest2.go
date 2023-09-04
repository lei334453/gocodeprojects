package main 
import "fmt"
//func 函数名 (形参列表)(返回值类型){
// 	函数体
// }
//函数体应该满足见名知意驼峰命名  
//参数可多可少  作用是接受外部参数  注意形参和实参的区别
//返回值类型什么都不写也可以
//注意返回值的个数
func sumadd(num1 int, num2 int)( int,int){
	var sum int = 0
	sum+=num1
	sum+=num2
	// return sum
	// fmt.Println(sum)

	var result int=num1-num2
	return sum,result
}
func main(){

	// sum1,result1:=sumadd(119,20)
	sum1,_:=sumadd(119,20)//可以忽略返回值
	fmt.Println(sum1)
}