package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'
	h := 'G'

	fmt.Printf("%T \n", a)  //int类型
	fmt.Printf("%T \n", b)  //string类型
	fmt.Printf("%T \n", c)  //float64类型
	fmt.Printf("%T \n", d)  //bool类型
	fmt.Printf("%T \n", e)  //string类型
	fmt.Printf("%T \n", f)  //string类型
	fmt.Printf("%T \n", g)  //int32类型
	fmt.Printf("%T \n", h)
}
