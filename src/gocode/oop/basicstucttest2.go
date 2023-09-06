package main
import "fmt"
type Teacher struct{
	Name string
	Age int
	School string 
}

func main(){
	//方式一
	// var t1 Teacher
	// t1.Name="赵闪闪"
	// t1.Age=31
	// t1.School="齐鲁大学"
   //方式二
    // var t Teacher=Teacher{"xiaolei",32,"清北科技大学"}
	// fmt.Println(t)
	//方式三  返回的是结构体指针
        //  var t *Teacher=new(Teacher)
		//  //t是指针 t其实指向的地址 应该给这个地址的指向的对象赋值
		//  (*t).Name="小磊"
		//  (*t).Age=34
		// //  (*t).School="亲隔壁五道口"
		// //简化了 为了符合程序员的编程习惯 简化写法
		// t.School="清北五道口"
		//  fmt.Println(*t)
     //第四种
	 var t*Teacher = &Teacher{}//这里也可以在括号里用前一种
	  (*t).Name="小磊"
	  (*t).Age=34
	  t.School="社会大学"
	  fmt.Println(*t)

		 
   


}