package main

import (
	"fmt"
	"time"
)

// soma entre dois números
var n1 int
var n2 int

func main() {
	fmt.Println("==== DESAFIO 03 ====")
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&n1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&n2)

	s := n1 + n2

	fmt.Printf("A soma de %v + %v equivale a %v", n1, n2, s)
	time.Sleep(10 * time.Second)
}