package main

import "fmt"

func main() {
    // Declare a slice
    var numbers []int

    // Append elements like you would push_back in C++
    numbers = append(numbers, 10)
    numbers = append(numbers, 20, 30)

    fmt.Println(numbers) // Output: [10 20 30]

    // Access elements
    fmt.Println(numbers[1]) // Output: 20

    // Iterate like a vector
	for i:=0;i<len(numbers);i++ {
		fmt.Printf("Index %d, Value %d\n", i, numbers[i])
	}
	// Alternatively, you can use a range loop
    // for i, v := range numbers {
    //     fmt.Printf("Index %d, Value %d\n", i, v)
    // }
}
