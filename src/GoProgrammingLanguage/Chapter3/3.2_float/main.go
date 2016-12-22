package main

import (
	"math"
	"fmt"
)

func main() {
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f + 1)    // "true"!
	var ff float32 = 1 << 24
	fmt.Println(ff)

	fmtfloat()

	fmt.Println(compute(2.2, 5.6, "*"))
	fmt.Println(compute(2.2, 0, "/"))
	fmt.Println(compute(22.2, 2.6, "/"))

}
/**
用Printf函数的%g参数打印浮点数，将采用更紧凑的表示形式打印，并提供足够的精度，但是对应表格的数据，
使用%e（带指数）或%f的形式打印可能更合适。

 */
func fmtfloat() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
		//打印精度是小数点后三个小数精度和8个字符宽度
	}
}

/**
如果一个函数返回的浮点数结果可能失败，最好的做法是用单独的标志报告失败，像这样：
 */
func compute(A float64, B float64, T string) (value float64, ok bool) {
	var result float64
	var failed bool
	switch T {
	case "*":
		result = A * B

		break
	case "/":
		if B != 0 {
			result = A / B
		} else {
			failed = true
		}


		break
	case "+":
		result = A + B

		break
	case "-":
		result = A - B

		break
	}

	if failed {
		return 0, false
	}
	return result, true
}