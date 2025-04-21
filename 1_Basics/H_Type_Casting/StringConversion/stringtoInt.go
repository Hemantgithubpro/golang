package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "42"
	v, _ := strconv.Atoi(s) // convert string to int

	fmt.Println(v) // 42
	fmt.Printf("type of v: %T\n", v) // int

	var i int = 42
	str := strconv.Itoa(i) // convert int to string

	fmt.Println(str) // 42
	fmt.Printf("type of str: %T\n", str) // string
}
