package main

import "fmt"

func main() {
	fmt.Println(multiply(3, 58))
	fmt.Printf("%T \n", multiply(2, 4))
}

func multiply(x int, y int) int {
	return x * y
}
