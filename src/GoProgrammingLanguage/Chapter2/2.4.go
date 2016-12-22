package main

import "fmt"

/**
元组赋值是另一种形式的赋值语句，它允许同时更新多个变量的值
 */
func main() {
	fmt.Println(gcd(12,20))
	fmt.Println(fib(9))
}
//计算两个整数值的的最大公约数（GCD）
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
//计算斐波纳契数列（Fibonacci）的第N个数
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}