## It is basically importing a package.
## Visit https://pkg.go.dev/std for standard library packages.

## Example for 'quote' package:
- Visit [quote](https://pkg.go.dev/rsc.io/quote/v4) package page.
- Write the code in `quote.go` file:
```go
package main
import (
    "fmt"
    "rsc.io/quote/v4"
)
func main() {
    fmt.Println(quote.Hello())
}
```
- Then to import the package, run the following command:
```bash
go mod init quote
go mod tidy
```
- To run the code, use the command:
```bash
go run quote.go
```