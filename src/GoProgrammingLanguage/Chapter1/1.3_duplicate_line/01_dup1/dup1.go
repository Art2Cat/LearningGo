package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
常用 转义字符
%d           int类型
%x, %o, %b   分别为16位，8位，2位的int类型
%f, %g, %e   浮点数表示类型
%t           布尔变量：true 或 fale
%c           rune(Unicode码点)， Go语言中特有的Unicode字符类型
%s           string
%q           带双引号的字符串“abc”或 带单引号的 rune 'c'
%v           会将任意变量以易读的形式打印出来
%T           打印变量的类型
%%           字符型百分比标志(表示%本身，没有其他操作)
*/

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}
	//注意忽略潜在的input.Err()错误
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
