package main

import (
	"fmt"
	"strconv"
)

//&& (and) || (or)
func main() {
	b := true
	fmt.Println(btoi(b))
	if b {
		if 1 == 1 || 0 != 0 {
			fmt.Printf("type is: %T, value is:%s \n", b, strconv.FormatBool(b))
		} else {
			fmt.Printf("type is: %T, value is:%s \n", b, strconv.FormatBool(b))
		}
	}
}

// btoi returns 1 if b is true and 0 if false.
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
