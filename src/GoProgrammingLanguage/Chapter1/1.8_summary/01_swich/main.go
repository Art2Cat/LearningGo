package main

import "fmt"

func main() {
	switchTest(signum(0))
}

/*在Go中switch语句是不需要break，如果你想要相邻的几个case都执行同一逻辑的话，
需要自己显式地写上一个fallthrough语句来覆盖这种默认行为
*/
func switchTest(i int) {
	switch i {
	case 0:
		fmt.Println("hello world")
	case 1:
		fmt.Println("Fuck world")
		fallthrough
	case 2:
		fmt.Println("fuck you too")
	}
}

//另外switch可以不用操作语句
func signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}
