package main

import (
	"fmt"
	"os"
)

/**
for 循环简介
*/
func main() {
	var k, j int
	//传统for循环
	for i := 0; i < 50; i++ {
		k += i
	}
	fmt.Println(k)

	//
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}

	//Go 是没有while循环的，可以使用for达成类似while循环
	for {
		j++
		fmt.Println(j)
		if j > 10 {
			return //这里使用break来跳出循环，也可以使用return
		}
	}

}
