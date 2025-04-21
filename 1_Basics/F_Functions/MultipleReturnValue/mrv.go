package main 
import "fmt"

func isDigit(c byte) bool {
    return c >= '0' && c <= '9'
}

func nextInt(b []byte, i int) (int, int) {
    for ; i < len(b) && !isDigit(b[i]); i++ {
    }
    x := 0
    for ; i < len(b) && isDigit(b[i]); i++ {
        x = x*10 + int(b[i]) - '0'
    }
    return x, i
}

func main() {
    b := []byte("123 456 789") // Define b as a byte slice
    for i, x := 0, 0; i < len(b); { // Declare x before the loop
        x, i = nextInt(b, i)
        fmt.Println(x)
    }
	fmt.Println("done")
}