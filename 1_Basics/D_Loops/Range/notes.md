## Normal range loop

```go
a := []int{1, 2, 3, 4, 5}
for i:= range a {
    fmt.Println(i, a[i])
}
```

## Range loop with index and value

- The range form of the for loop iterates over a slice or map.
- When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
- it will first return the index and then the value of the element at that index.
```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
for i, v := range pow {
    fmt.Printf("2**%d = %d\n", i, v)
}
```

## Range loop with only value
- If you only want the value, you can use the blank identifier `_` to ignore the index.
```go
for _, v := range pow {
    fmt.Printf("%d\n", v)
}
```
## Range loop with only index
- If you only want the index, you can use the blank identifier `_` to ignore the value.
```go
for i, _ := range pow {
    fmt.Printf("%d\n", i)
}
```