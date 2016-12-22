package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	a := "a/b/c.go"
	b := "c.d.go"
	c := "abc"
	f := "1234567890.2235"
	d := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '.', '2', '2', '3', '5'}
	fmt.Println(basename(a))                  // "c"
	fmt.Println(basename(b))                  // "c.d"
	fmt.Println(basename(c))                  // "abc"
	fmt.Println(basename2(a))                 // "c"
	fmt.Println(basename2(b))                 // "c.d"
	fmt.Println(basename2(c))                 // "abc"
	fmt.Println(comma(f))                     // a/,b/c,.go
	fmt.Println(commabyte(d))                 // a/,b/c,.go
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		return comma(s[:dot-3]) + "," + s[dot-3:]
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//3.10 solution
func commabyte(b []byte) string {
	var y string
	n := len(b)
	var buf bytes.Buffer
	if n <= 3 {
		return string(b)
	}

	for index := 0; index < n; index++ {
		v := b[index]
		if n%3 == 0 {
			if index == 3 {
				buf.WriteByte(',')
				buf.WriteByte(v)
			} else if (index-3)%3 == 0 && index-3 > 0 {
				buf.WriteByte(',')
				buf.WriteByte(v)
			} else {
				buf.WriteByte(v)
			}
		} else if n%3 == 1 {
			if index == 1 {
				buf.WriteByte(',')
				buf.WriteByte(v)
			} else if (index-1)%3 == 0 && index-1 > 0 {
				buf.WriteByte(',')
				buf.WriteByte(v)
			} else {
				buf.WriteByte(v)
			}
		} else if n%3 == 2 {
			if index == 2 {
				buf.WriteByte(',')
				buf.WriteByte(v)
			} else if (index-2)%3 == 0 && index-2 > 0 {
				buf.WriteByte(',')
				buf.WriteByte(v)
			} else {
				buf.WriteByte(v)
			}
		}
	}
	y = buf.String()
	return y
}

// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
