package main

import "fmt"

func main() {
	i := 5
	if i < 10 {
		fmt.Println(i)
		fmt.Println(string(256))
	}
}

func init_statement_if() {
	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

}
