// we mostly use arrays in the form of slices, its like a dynamic array(vector in c++)
package main 
import "fmt"
func main() {
	// var arr [6]int;
	// arr[0] = 1
	// arr[1] = 2
	// arr[2] = 3
	// arr[3] = 4
	// arr[4] = 5
	// arr[5] = 6

	// for i:= range arr {
	// 	// Print the value of the array at index i
	// 	fmt.Println(arr[i])
	// }

	arr2:= [5]string{"john", "doe", "jane", "smith", "bob"}
	for i:= range arr2 {
		fmt.Println(arr2[i])
	}
	
}