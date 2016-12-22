package main

import "fmt"

var s string = "hello, world"

func main() {
	indexString()

	fmt.Println("goodbye" + s[5:])
}

func indexString() {
	fmt.Println(len(s))     // "12"
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
}
