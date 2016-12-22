package main

import "fmt"

var data string = "Hello, 世界"

// /n 换行 /t 制表符 /v 垂直制表符 /b 退格 /f 换页 /r 回车
// 十六进制的转义形式是\xhh，其中两个h表示十六进制数字（大写或小写都可以）
// 八进制转义形式是\ooo，包含三个八进制的o数字（0到7），但是不能超过\377（译注：对应一个字节的范围，十进制为255）
func main() {
	for i, r := range data {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	// "program" in Japanese katakana
	s := "プログラム"
	fmt.Println(s)
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"
}
