Earlier I used:
```go
package main
println("hello world")
```
and it worked. Why did it work without importing `fmt`?

In Go, the `println` function is a built-in function provided by the Go runtime, so it does not require importing the `fmt` package. Built-in functions like `println` and `print` are always available without any imports.

However, `println` is intended primarily for debugging purposes and is not as flexible or feature-rich as `fmt.Println`. For example:

1. **`println` is not formatted output**: It simply prints the arguments separated by spaces, without any formatting options.
2. **`println` is not guaranteed to behave consistently across platforms**: It is not part of the Go standard library and is considered low-level.
3. **`fmt.Println` is preferred for production code**: It provides more control and is part of the `fmt` package, which is designed for formatted I/O.

If you want to use `fmt.Println`, you would need to uncomment the `import "fmt"` line and replace `println` with `fmt.Println`. Here's how it would look:

```go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

This is the idiomatic way to print output in Go.