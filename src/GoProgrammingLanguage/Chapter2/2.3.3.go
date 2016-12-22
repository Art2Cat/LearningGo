package main

import "fmt"
/**
创建变量的另一个方法是调用用内建的new函数
new函数使用常见相对比较少，因为对应结构体来说，
可以直接用字面量语法创建新变量的方法会更灵活。
 */
func main() {
	p := new(int)   // p, *int 类型, 指向匿名的 int 变量
	fmt.Println(*p) // "0"
	*p = 2          // 设置 int 匿名变量的值为 2
	fmt.Println(*p) // "2"

	//每次调用new函数都是返回一个新的变量的地址，因此下面两个地址是不同的：
	q := new(int)
	s := new(int)
	fmt.Println(q == s) // "false"
}
