## Variables in Go
- Variables are used to store data in a program.
- In Go, variables are declared using the `var` keyword or the short variable declaration syntax `:=`.
- Constants are declared using the `const` keyword.
- Variables can be of different types, such as `int`, `float64`, `string`, etc.

## Key points about := shorthand declaration:

- Only works inside functions
- Can't be used for package-level variables
- Automatically infers the type based on the value
- Can't be used to declare constants
- Can declare multiple variables in one line: `x, y := 10, 20`
- Can't redeclare variables using `:=` unless at least one new variable is introduced