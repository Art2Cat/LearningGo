package main

import "fmt"

func main() {
	CInt8()
	test()
	conInt()
	fmtInt()
}

/**
如果一个算术运算的结果，不管是有符号或者是无符号的，如果需要更多的bit位才能正确表示的话，就说明计算结果是溢出了。
超出的高位的bit位部分将被丢弃。如果原始的数值是有符号类型，而且最左边的bit为是1的话，那么最终结果可能是负的，
例如int8的例子：
 */

func CInt8() {
	var u uint8 = 255
	fmt.Println(u, u + 1, u * u, u + u)
	var i int8 = 127
	fmt.Println(i, i + 1, i * i, i + i)
}

/**
位操作运算符^作为二元运算符时是按位异或（XOR），当用作一元运算符时表示按位取反；也就是说，
它返回一个每个bit位都取反的数。位操作运算符&^用于按位置零（AND NOT）：表达式z = x &^ y结果z的bit位为0，
如果对应y中bit位为1的话，否则对应的bit位等于x相应的bit位的值。

无符号数往往只有在位运算或其它特殊的运算场景才会使用，就像bit集合、分析二进制文件格式或者是哈希和加密操作等。
 */
func test() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1} x与y共有1
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5} x与y集合1，2，5
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}  x与y 对称差
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}

func conInt() {
	var apples int32 = 1
	var oranges int16 = 2
	var compote = int(apples) + int(oranges)
	fmt.Printf("compote equals: %d\n", compote)
}

/**
请注意fmt的两个使用技巧。通常Printf格式化字符串包含多个%参数时将会包含对应相同数量的额外操作数，
但是%之后的[1]副词告诉Printf函数再次使用第一个操作数。第二，%后的#副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀。
字符使用%c参数打印，或者是用%q参数打印带单引号的字符：
 */
func fmtInt()  {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
}