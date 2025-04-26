package main

import (
    "fmt"
)

func main() {
    // Create an unbuffered channel for int values
    ch := make(chan int)

    // Start a goroutine that sends a value
    go func() {
        ch <- 42 // Send the value 42
    }()

    // Receive the value from the channel
    value := <-ch
    fmt.Println("Received:", value)
}