# Channels in golang
Channels are typed conduits for communication between goroutines. They allow safe data transmission between concurrent processes.

## Channel Basics
- Created using `make(chan Type)`
- Send values with `ch <- v`
- Receive values with `v := <-ch`
- Close channels using `close(ch)`

## Channel Types
1. **Unbuffered Channels**
    - Synchronous communication
    - Sender blocks until receiver is ready
    - `ch := make(chan int)`

2. **Buffered Channels**
    - Asynchronous up to buffer size
    - `ch := make(chan int, 100)`

## Key Features
- Thread-safe by design
- First-in-first-out (FIFO)
- Can be ranged over
- Support select statements

## Common Patterns
1. **Generator Pattern**
    - Function returns channel
    - Produces values asynchronously

2. **Pipeline Pattern**
    - Chain of goroutines
    - Each stage connects via channels

## Best Practices
- Always close sender side
- Check for closed channels
- Use directional channels when possible
- Handle deadlocks carefully

## Examples

1. **Basic Channel Usage**
```go
func main() {
    ch := make(chan string)
    go func() {
        ch <- "Hello from goroutine!"
    }()
    msg := <-ch
    fmt.Println(msg)
}
```

2. **Buffered Channel**
```go
func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

3. **Range Over Channel**
```go
func main() {
    ch := make(chan int)
    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
        }
        close(ch)
    }()
    
    for num := range ch {
        fmt.Println(num)
    }
}
```

4. **Select Statement**
```go
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() { ch1 <- "from 1" }()
    go func() { ch2 <- "from 2" }()

    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    }
}
```

5. **Channel Direction**
```go
func send(ch chan<- int) {
    ch <- 42
}

func receive(ch <-chan int) {
    val := <-ch
    fmt.Println(val)
}

func main() {
    ch := make(chan int)
    go send(ch)
    receive(ch)
}
```