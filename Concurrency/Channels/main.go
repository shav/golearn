package main

import (
    "fmt"
    "strings"
    "time"
)

func main() {
    // Создание канала
    msg1 := make(chan string, 2 /* capacity for buffered channel */)
    fmt.Printf("len(msg1): %d, cap(msg1): %d\n", len(msg1), cap(msg1))
    msg2 := make(chan string)
    fmt.Printf("len(msg2): %d, cap(msg2): %d\n", len(msg2), cap(msg2))
    
    go capitalize("Hello", msg1)
    go capitalize("World", msg2)
    
    // Чтение из канала
    fmt.Println(<-msg1)
    fmt.Println(<-msg2)
    
    go func(msg string) {
        fmt.Println(msg)
    }("going")
    
    fmt.Println("---------------------------------------------------")
    
    c1 := make(chan string)
    c2 := make(chan string)
    signals := make(chan bool)
    
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        close(c2)
    }()
    
    // Switch-case запись в каналы
    select {
    case signals <- true:
        fmt.Println("sent signal: true")
    // Если канал заблокирован:
    default:
        fmt.Println("no signals sent")
    }
    
    // Switch-case чтение из каналов
    for i := 0; i <= 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received ", msg1)
        case msg2, isOpened := <-c2:
            if isOpened {
                fmt.Println("received ", msg2)
            } else {
                fmt.Println("msg2 is closed")
            }
            // Если все каналы пусты:
            // default:
            //  fmt.Println("no message activity")
        }
    }
    
    fmt.Println("---------------------------------------------------")
    
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)
    
    // Чтение из каналов через foreach
    for item := range queue {
        fmt.Println(item)
    }
}

func capitalize(name string, result chan<- string) {
    for i := 0; i < 3; i++ {
        fmt.Println(name, ":", i)
    }
    // Запись в канал
    result <- strings.ToTitle(name)
}
