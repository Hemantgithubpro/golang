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