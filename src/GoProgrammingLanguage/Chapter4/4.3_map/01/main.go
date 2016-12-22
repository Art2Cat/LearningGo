package main

import "fmt"

func main() {
	// 创建map方法一
	maps := make(map[string]int)
	maps["monday"] = 1
	maps["tuesday"] = 2
	maps["wednesday"] = 3
	maps["thursday"] = 4
	// 创建map方法一
	mps := map[string]int{"first": 1, "second": 2, "third": 3}
	// map循环取出所有key和value
	for key, value := range mps {
		fmt.Printf("%s is %d\n", key, value)
	}

	fmt.Println(len(maps))
	// 删除指定值
	delete(maps, "wednesday")
	fmt.Println(maps)
}
