package main

import "fmt"

func main() {
	// for i := 1; i < 5; i++ {
	// 	fmt.Println("Hello World", i)
	// }

	for i := range 5 {
		fmt.Printf("Hello World->%d\n", i)
	}
}
