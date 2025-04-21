package main
import "fmt"
func switchCase(){
	x := 10
	switch {
	case x > 5:
		fmt.Println("x is greater than 5")
	case x == 5:
		fmt.Println("x is equal to 5")
	default:
		fmt.Println("x is less than 5")
	}
}
func main(){
	switchCase()
	
}