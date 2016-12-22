package main

import "fmt"
//常量boilingF是在包一级范围声明语句声明的，可在整个包一级范围内访问
const boilingF = 212.0

func main() {
	//f和c两个变量是在main函数内部声明的声明语句声明的，只能在函数内部访问
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C
}
