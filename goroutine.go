package main

import (
    "fmt"
    "runtime"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        runtime.Gosched()   // await, 让 CPU 把时间片给别人
        fmt.Println(s)
    }
}


// channel 共享数据
func sum(a []int, c chan int) {
    total := 0
    for _, v := range a {
        total += v
    }
    c <- total
}

func fibonacci(n int, c chan int) {
    x, y := 1, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func sfibonacci(c, quit chan int) {
    x, y := 1, 1
    for {
        select {
            case c <- x:
                x, y = y, x+y
            case <-quit:
                fmt.Println("quit")
                return
        }
    }
}

func main() {
    go say("world")
    say("Hello")

    a := []int{1, 2, 3, 4, 5}
    // 无缓冲 channel 接受和发送数据都是 block 的
    c := make(chan int)     // channel 只能用 make 来创建
    // 缓冲 channel（buffered channels)
    // c := make(chan int, value)       value 大于0时，channel 是有缓冲，无阻塞的,写满 value 个数据才开始阻塞写入

    go sum(a[:len(a)/2], c)
    go sum(a[len(a)/2:], c)
    x, y := <-c, <-c

    fmt.Println(x, y, x+y)

    e := make(chan int, 3)
    e <- 1
    e <- 2
    fmt.Println(<-e)
    fmt.Println(<-e)

    f := make(chan int, 10)
    go fibonacci(cap(f), f)
    for i := range f {
        fmt.Println(i)
    }

    s := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println("s")
            fmt.Println(<-s)
        }
        quit <- 0
    }()
    sfibonacci(s, quit)


    // timeout
    go func() {
        for {
            select {
                case v := <- s:
                    fmt.Println("sssss", v)
                case <- time.After(5 * time.Second):
                    fmt.Println("timeout")
                    quit <- 1
                    break
            }
        }
    }()
    <- quit

    fmt.Println(runtime.NumCPU(), runtime.NumGoroutine())

}
