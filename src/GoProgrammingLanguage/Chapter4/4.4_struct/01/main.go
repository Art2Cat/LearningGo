package main

import "fmt"

type Language struct {
	title   string
	name    string
	website string
	id      int
}

func main() {
	var Langugage1 Language /* 声明 Langugage1 为 Language 类型 */
	var Langugage2 Language /* 声明 Langugage2 为 Language 类型 */

	/* Langugage 1 描述 */
	Langugage1.title = "Go官网"
	Langugage1.website = "www.golang.org"
	Langugage1.name = "Go"
	Langugage1.id = 00001

	/* Language 2 描述 */
	Langugage2.title = "Python官网"
	Langugage2.website = "www.python.org"
	Langugage2.name = "Python"
	Langugage2.id = 00002

	/* 打印 Language1 信息 */
	fmt.Printf("Langugage 1 title : %s\n", Langugage1.title)
	fmt.Printf("Langugage 1 website : %s\n", Langugage1.website)
	fmt.Printf("Langugage 1 name : %s\n", Langugage1.name)
	fmt.Printf("Langugage 1 id : %d\n", Langugage1.id)

	/* 打印 Language2 信息 */
	fmt.Printf("Langugage 2 title : %s\n", Langugage2.title)
	fmt.Printf("Langugage 2 website : %s\n", Langugage2.website)
	fmt.Printf("Langugage 2 name : %s\n", Langugage2.name)
	fmt.Printf("Langugage 2 id : %d\n", Langugage2.id)
}
