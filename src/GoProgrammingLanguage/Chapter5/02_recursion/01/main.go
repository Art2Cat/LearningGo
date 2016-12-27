// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
    "fmt"
    "os"
    "./html"

)

func main() {
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
        os.Exit(1)
    }
    for _, link := range visit(nil, doc) {
        fmt.Println(link)
    }
}
func visit(i interface{}, node *html.Node) (interface{}, *html.Node) {
	return i, node
}
