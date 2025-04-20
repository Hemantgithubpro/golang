package main

import (
	"fmt"
	// "math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<63
	// z      complex128 = cmplx.Sqrt(-5 + 12i)
	z      complex128 = 5 + 12i // complex number (real + imaginary part)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// }

// func main(){
// 	a := Person{name: "John", age: 30}
// 	fmt.Println("Name:", a.name)
// 	fmt.Println("Age:", a.age)
// 	fmt.Println("Person:", a)
// 	fmt.Printf("Type of a: %T\n", a)
// 	fmt.Printf("Hello, %s! You are %d years old.\n", a.name, a.age)
// 	fmt.Printf("Hello, %v! You are %v years old.\n", a.name, a.age)
// }
