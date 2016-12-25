package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// 这里按照书上的os.Args[1:] 打印不出命令行参数，改成[0:]就可以打印
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = ""
	}

	fmt.Println(s)
}
