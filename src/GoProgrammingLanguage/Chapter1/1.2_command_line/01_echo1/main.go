package main

import (
	"fmt"
	"os"
)

/**
打印命令行参数
*/
func main() {
	var s, sep string
	fmt.Println(os.Args)
	fmt.Println(len(os.Args)) //为嘛长度为1？
	for index := 0; index < len(os.Args); index++ {
		s += sep + os.Args[index]
		sep = " "
	}
	fmt.Println(s)
}
