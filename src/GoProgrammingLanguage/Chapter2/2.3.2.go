package main

import "fmt"
/**
一个指针对应变量在内存中的存储位置
指针被称之为“指向int类型的指针”

 */
var p = f()

func main() {
	x := 1
	p := &x         // p, 指针对应的数据类型是*int, 指向 x，或者p指针保存了x变量的内存地址
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"

	//任何类型的指针的零值都是nil。如果p != nil测试为真，那么p是指向某个有效变量。
	//指针之间也是可以进行相等测试的，只有当它们指向同一个变量或全部是nil时才相等。
	pointerC()

	//在Go语言中，返回函数中局部变量的地址也是安全的。
	// 例如下面的代码，调用f函数时创建局部变量v，在局部变量地址被返回之后依然有效，因为指针p依然引用这个变量。
	fmt.Println(f() == f()) // "false"

	/**
	因为指针包含了一个变量的地址，因此如果将指针作为参数调用函数，那将可以在函数中通过该指针来更新变量的值。
	例如下面这个例子就是通过指针来更新变量的值，然后返回更新后的值，可用在一个表达式中
	（译注：这是对C语言中++v操作的模拟，这里只是为了说明指针的用法，incr函数模拟的做法并不推荐）：
	 */
	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)
}

func pointerC() {
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}