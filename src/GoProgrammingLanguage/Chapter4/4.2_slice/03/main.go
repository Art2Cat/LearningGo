package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	fmt.Println(rotate(s))
	fmt.Println(rotate(s))
}

func reverse(s []int) {
	y := len(s)
	var x [6]*int
	for i, j := 0, y-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = &s[j], &s[i]
		s[i], s[j] = *x[i], *x[j]
	}
}

func rotate(s []int) []int {
	x := 2
	for y := 0; y < x; y++ {
		reverse(s)
	}
	return s
}
