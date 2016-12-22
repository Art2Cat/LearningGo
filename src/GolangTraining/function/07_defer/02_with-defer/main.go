package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world() //defer关键字 将执行函数延期
	hello()
	//print hello world
}
