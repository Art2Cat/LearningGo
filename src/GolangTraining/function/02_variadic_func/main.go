package main

import "fmt"

func main() {
	s := []float64{1, 23, 36, 4, 5}
	fmt.Println(average(1, 1, 50, 6 ,9))
	fmt.Println(average(s...))
	fmt.Println(averageS(s))
}

func average(f ...float64) float64 {

	fmt.Printf("%T \n ", f)
	var total float64
	for _, x := range f {
		total += x
	}

	return total / float64(len(f))
}

func averageS(f []float64) float64 {

	fmt.Printf("%T \n ", f)
	var total float64
	for _, x := range f {
		total += x
	}

	return total / float64(len(f))
}