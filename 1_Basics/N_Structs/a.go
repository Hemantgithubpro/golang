package main
import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person // Embedding Person struct
	Salary int
}

func main() {
	// Create an instance of Employee
	emp := Employee{
		Person: Person{Name: "Alice", Age: 30},
		Salary: 50000,
	}

	// Access fields from the embedded Person struct
	fmt.Println("Name:", emp.Name) // Accessing Name from Person
	fmt.Println("Age:", emp.Age)   // Accessing Age from Person
	fmt.Println("Salary:", emp.Salary)
}