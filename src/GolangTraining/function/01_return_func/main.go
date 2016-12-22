package main

import "fmt"

func main() {
	fmt.Println(hello("yiming"))
	fmt.Println(helloT("yiming", "huang"))
}

func hello(name string) string {
	s := fmt.Sprint("Hello ", name)
	return s
}

func helloT(fname, lname string) (string, string) {
	s := fmt.Sprint("Hello ", fname, " ", lname)
	n := fmt.Sprint("Hello ", lname, fname)
	return s, n
}