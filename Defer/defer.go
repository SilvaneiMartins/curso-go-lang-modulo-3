package main

import "fmt"

func dia1() {
	fmt.Println("Segunda Feira")
}

func dia2() {
	fmt.Println("Terça Feira")
}

func main() {
	defer dia2()
	dia1()
}
