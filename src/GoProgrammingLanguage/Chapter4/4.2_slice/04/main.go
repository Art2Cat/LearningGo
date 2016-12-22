package main

import "fmt"

func main() {
	s := []string{"a", "b", "b", "b", "b", "c", "d", "d", "e", "f"}
	fmt.Println(s)
	fmt.Println(deleteDuplicate(s))
}

func deleteDuplicate(s []string) []string {
	var st string
	var flag bool
	for {
		if flag {
			break
		}
		for index := 1; index < len(s); index++ {
			st += s[index]
			if index+1 != len(s) && s[index] == s[index-1] {
				s = remove(s, index)
				return deleteDuplicate(s)
			} else {
				flag = true
			}
		}

	}
	return s
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
