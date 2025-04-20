package main
import "fmt"
import "math"
func main(){
	// var arr = []int{1, 2, 3, 4, 5,45,12}
	// for i:=0; i<5; i++{
	// 	fmt.Println("Array elements:", arr[i])
	// }
	// for i:= range arr{
	// 	fmt.Println("Array elements:", arr[i])
	// }
	// fmt.Println("array length:", len(arr))

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}


	a := math.Pow(3, 2)
	fmt.Println("Value of a:", a)
}