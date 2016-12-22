package main

import "fmt"

func main() {
	var x [50]string
	for index := 0; index < 50; index++ {
		x[index] = string(index)
	}

	fmt.Printf("%T\n %d\n", x, len(x))
}
