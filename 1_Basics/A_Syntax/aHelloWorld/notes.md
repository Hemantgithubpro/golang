# Steps for Hello World
1. create `hello.go` file
2. enable dependency tracking for my code: `go mod init hello` or `go mod init <module_name>` or `go mod init example/hello` in terminal.
3. write code in `hello.go` file:
```go
package main
import "fmt"
func main() {
    fmt.Println("Hello, World!")
}
```
4. run the code: `go run hello.go` in terminal.
