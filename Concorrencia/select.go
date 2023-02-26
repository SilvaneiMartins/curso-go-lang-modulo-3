package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "Um"
    }()

    go func() {
        time.Sleep(time.Second * 3)
        c2 <- "Dois"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 :=  <-c1:
            fmt.Println("Mensagem do canal 1:", msg1)
        case msg2 := <-c2:
            fmt.Println("Mensagem do canal 2:", msg2)
        }
    }
}
