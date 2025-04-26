package main
import "fmt"
func make1(){
	// make() is a built-in function that allocates and initializes slices, maps, and channels.
	
	// Slices
	s:=make([]int, 5) // creates a slice of int with length 5 and capacity 5
	fmt.Println("Slice:", s) // Output: Slice: [0 0 0 0 0]
	s = append(s, 1, 2, 3) // appending elements to the slice
	fmt.Println("Slice after append:", s) // Output: Slice after append: [0 0 0 0 0 1 2 3]
	
}