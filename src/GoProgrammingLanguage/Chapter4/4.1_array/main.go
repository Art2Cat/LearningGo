package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var a [3]int = [3]int{1, 2, 3} // array of 3 integers
	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	//数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始化值的个数来计算
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
