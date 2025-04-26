package main
import "fmt"
func make2(){
	// make() is a built-in function that allocates and initializes slices, maps, and channels.
	
	// Maps
	m:=make(map[string]int, 5) // creates a map with string keys and int values, with an initial capacity of 5
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	// m["four"] = 4
	
	fmt.Println("Map:", m) // Output: Map: map[one:1 three:3 two:2]
	
}