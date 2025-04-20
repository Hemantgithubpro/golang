package main
import "fmt"

type Person struct {
	name string
	age  int
}

func main(){
	a := Person{name: "John", age: 30}
	fmt.Println("Name:", a.name)
	fmt.Println("Age:", a.age)
	fmt.Println("Person:", a)
	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Hello, %s! You are %d years old.\n", a.name, a.age)
	fmt.Printf("Hello, %v! You are %v years old.\n", a.name, a.age)
}