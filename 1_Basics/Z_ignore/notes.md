I studied from this folder: https://go.dev/doc/tutorial/getting-started


# Call code in an external package
When you need your code to do something that might have been implemented by someone else, you can look for a package that has functions you can use in your code.

# For example
if you want to use a package that can make HTTP requests, you can use the `net/http` package. You can find the documentation for this package at https://pkg.go.dev/net/http.
```go
package main
import (
    "fmt"
    "net/http"
)
func main() {
    // Make an HTTP GET request to a URL
    response, err := http.Get("https://example.com")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer response.Body.Close()

    // Print the status code of the response
    fmt.Println("Response Status Code:", response.StatusCode)
}
```
### In this example
we import the `net/http` package and use its `Get` function to make an HTTP GET request to a specified URL. The response is then printed to the console.

## one more thing 
- you can run: `go build main.go` to build the code and create an executable file. This file can be run directly from the command line, and it will execute the code in `main.go`.
- This is useful for creating standalone applications that can be distributed and run on different systems without requiring the Go runtime to be installed.