package main
import "fmt"

func main(){
	var score int = 87
    switch score/10{
	case 10:
		fmt.Println("您的等级为A级,你很有完美~")
	case 9:
		fmt.Println("您的等级为B级,你很优秀~")
	case 8:
		fmt.Println("您的等级为C级,你优秀~")
	case 7:
		fmt.Println("您的等级为D级,你还有上升空间~")
	case 6:
		fmt.Println("您的等级为E级,你太差了，需要努力~")
	case 5:
		fmt.Println("您的等级为F级,你的进步空间太大,你需要加油加油加油~")
	default: fmt.Println("你的成绩有误")//可以放在任意位置 可有可无
	}
	//奇怪的用法 跟if类似
	switch {
	case score==87 :
		fmt.Println("aaa")
	case score==89 :
		fmt.Println("bbb")
	}
}