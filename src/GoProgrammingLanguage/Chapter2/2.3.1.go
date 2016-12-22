package main

import "fmt"
/**
变量声明的一般语法如下
var 变量名字 类型 = 表达式
var varName type = expression
其中“类型”或“= 表达式”两个部分可以省略其中的一个
简短变量声明
“变量名字 := 表达式”

 */
func main() {
	//这里使用简短变量声明，简短变量声明对已声明过的变量只有赋值行为,如下代码所示
	n1 := 1
	n2, n1 := 1, 3
	n3, n2 := 1, 2
	fmt.Printf("n1 = %d\n" , n1)  //n1 = 3
	fmt.Printf("n2 = %d\n" , n2)  //n2 = 2
	fmt.Printf("n3 = %d" , n3)  //n3 = 1
}