package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 这里按照书上的os.Args[1:] 打印不出命令行参数，改成[0:]就可以打印
	fmt.Println(strings.Join(os.Args[0:], ""))
}
