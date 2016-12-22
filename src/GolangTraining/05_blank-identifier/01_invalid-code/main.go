package main

import "fmt"

func main() {
	a := "stored in a"
	b := "stored in b"
	fmt.Println("a - ", a)
	// 变量b并没有被使用，编译器会报错
}
