package main

import (
	"fmt"
	"os"
)

/**
这个练习有点无解，长度才1
*/
func main() {
	// 这里按照书上的os.Args[1:] 打印不出命令行参数，改成[0:]就可以打印
	for x, arg := range os.Args[0:] {
		fmt.Println(x)
		fmt.Println(arg)
	}
}
