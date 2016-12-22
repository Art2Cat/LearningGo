// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

func main() {
	//这里使用在main函数中定义的两个常量作为被调用fToC函数的参数
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
}
//这里函数fToC的f参数是float64类型，函数运行后返回结果为float64类型的值
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}