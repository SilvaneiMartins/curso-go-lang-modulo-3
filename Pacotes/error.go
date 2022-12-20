package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v: %s", e.When, e.What)
}

func oops() error {
	return &MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"O arquivo do sistema não foi encontrado!",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
