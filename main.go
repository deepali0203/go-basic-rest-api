package main

import (
	"fmt"
)

func main() {
	c := 42
	r := "hello"
	fmt.Printf("%s_%v", r, c)
	fmt.Println("xyz")
	s := fmt.Sprintf("dwivedi %s", "go mabc")
	fmt.Printf("shubham" + s)
}
