package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a sample file with some content
	fileName := "example.txt"
	content := []byte("Hello, this is a sample file content.")
	err := os.WriteFile(fileName, content, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// defer os.Remove(fileName) // Clean up after execution

	// Open the file for reading
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read content from the file
	buffer := make([]byte, 16) // Read in chunks of 16 bytes
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Print(string(buffer[:n]))
	}
}