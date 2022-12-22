package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	h := sha1.New()

	h.Write([]byte("Silvanei Martins"))
	v := h.Sum([]byte{})
	fmt.Println(v)
}
