package main

import (
	"encoding/json"
	"os"
	"strings"
	"fmt"
)

func main() {
	cat := Cat{"meow", "yellow", 2,}
	bs := encode(cat)
	fmt.Println(&cat.Meow)
	fmt.Println(&cat.Color)
	fmt.Println(&cat.Age)
	fmt.Println(cat)
	fmt.Printf("%T \n", cat)


	cat = decode(bs)
	fmt.Println(&cat.Meow)
	fmt.Println(&cat.Color)
	fmt.Println(&cat.Age)
	fmt.Println(cat)
	fmt.Printf("%T \n", cat)
}

type Cat struct {
	Meow  string
	Color string
	Age   int
}

func encode(cat Cat) []byte {
	json.NewEncoder(os.Stdout).Encode(&cat)
	bs, _ := json.Marshal(cat)
	return bs
}

func decode(b []byte) Cat {
	var cat Cat
	json.NewDecoder(strings.NewReader(string(b))).Decode(&cat)
	return cat
}
