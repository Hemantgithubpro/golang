package main
import "fmt"
func prac1(){
	// create a slice of integers with 5 elements
	slice1:= []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	// printintarray(slice1)

	// var slice2 []string = []string{"burger", "pizza", "pasta"}
	// printstringarray(slice2)
}

func printintarray(arr []int) {
	for i:=range arr{
		fmt.Println(arr[i])
	}
}

func printstringarray(arr []string){
	for i:=range arr{
		fmt.Println(arr[i])
	}
}

func prac2(){
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
}

func prac3(){
	slice:=[]int{}
	slice=append(slice, 1,2,3,4,5)
	fmt.Println(slice)
	slice = append(slice, 4*3)
	fmt.Println(slice)
}

func prac4(){
	// create a slice from an array
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4] // slice from index 1 to 3 (4 is excluded)
	fmt.Println(slice) // Output: [2 3 4]
}

func prac5(){
	// create a slice with make function
	slice := make([]int, 5) // create a slice of length 5
	fmt.Println(slice) // Output: [0 0 0 0 0]
	slice[0] = 1 // Output: [1 0 0 0 0]

	slice2:= make([]int, 5, 10) // create a slice of length 5 and capacity 10
	fmt.Println(slice2) // Output: [0 0 0 0 0]
}