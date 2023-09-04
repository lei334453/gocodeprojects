package main
import "fmt"

func main(){
	var  arr[2][3]int16
	fmt.Println(arr)

	fmt.Printf("arr的地址值是:%p",&arr)
	fmt.Printf("arr[0]的地址值是:%p",&arr[0])
	fmt.Printf("arr[0][0]的地址值是:%p",&arr[0][0])

	fmt.Printf("arr[1]的地址值是:%p",&arr[1])
	fmt.Printf("arr[0][0]的地址值是:%p",&arr[1][0])
	arr[0][1]=47
	arr[0][0]=83
	arr[1][1]=90

    fmt.Println(arr)
	var arr1[2][3]int=[2][3]int{{1,4,5},{2,3,4}}
	fmt.Println(arr1)

	//二维数组的遍历

	var arr6[3][3]int=[3][3]int{{1,7,9},{2,5,8},{3,6,9}}
	fmt.Println(arr6)

	fmt.Println("===============================================")
	for i:=0;i<len(arr6);i++{
		for j:=0;j<len(arr6[i]);j++{
			fmt.Println(arr6[i][j],"\t")
		}
		fmt.Println()
	}

	fmt.Println("===============================================")
	//方式二
    for key,value :=range arr6{
		for k,v :=range value{
            fmt.Printf("arr6[%v][%v]=%v\t",key,k,v)
		}
		fmt.Println()
	}

}