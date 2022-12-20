package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if _, err := io.WriteString(os.Stderr, "Silvanei Martins!"); err != nil {
		log.Fatal(err)
	}
}
