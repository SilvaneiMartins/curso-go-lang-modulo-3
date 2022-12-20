package main

import (
	"io/ioutil"
	"log"
)

func main() {
	message := []byte("Silvanei Martins!\n")

	err := ioutil.WriteFile("Desenvolvedor", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
