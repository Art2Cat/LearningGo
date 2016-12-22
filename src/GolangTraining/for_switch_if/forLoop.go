package main

import "fmt"

func main() {
	//normalForLoop()
	//nestedForLoop()
	//whileForLoop()
	infiniteForLoop()
}

func normalForLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func nestedForLoop() {
	for i := 0; i < 10; i++ {
		for k := 0; k < 5; k++ {
			fmt.Println(i, "-", k)
		}
	}
}

func whileForLoop() {
	var i int
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func infiniteForLoop() {
	var i int
	for {
		fmt.Println(i)
		i++
		if i >= 15 {
			break
		}
	}
}