package main 
import (
	"fmt"
	"strings"
)

func main(){
	str:="Hello World"
	fmt.Println(str)
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.LastIndex(str,"l"))
}