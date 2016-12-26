package main

import "fmt"

/**
变量的生命周期指的是在程序运行期间变量有效存在的时间间隔。
对于在包一级声明的变量来说，它们的生命周期和整个程序的运行周期是一致的。
而相比之下，在局部变量的声明周期则是动态的：从每次创建一个新变量的声明语句开始，直到该变量不再被引用为止，
然后变量的存储空间可能被回收。函数的参数变量和返回值变量都是局部变量。它们在函数每次被调用的时候创建。
 */

var global *int

func x() {
	var x int
	x = 1
	global = &x
	fmt.Println(global)
}

func g() {
	y := new(int)
	*y = 1
	fmt.Println(y)
}
func main() {
	x()
	g()
}
