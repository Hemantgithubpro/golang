package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]	// slice of primes from index 1 to 4 (exclusive) or 1 to 3(inclusive)
	fmt.Println(s)	// [3 5 7]
}
