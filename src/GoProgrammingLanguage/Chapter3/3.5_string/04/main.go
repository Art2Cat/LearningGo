package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Enter numbers: ")
	fmt.Scan(&s)

	fmt.Println(comma(s))

}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if strings.Contains(s, "-") {
		if n-1 <= 3 {
			return s
		}

		return comma(s[:n-4]) + "," + s[n-3:]
	} else {
		if n <= 3 {
			return s
		}

		return comma(s[:n-3]) + "," + s[n-3:]
	}
	//if strings.Contains(s, ".") {
	//dot := strings.LastIndex(s, ".")
	//}
	//return s
}
