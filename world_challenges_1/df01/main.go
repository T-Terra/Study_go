package main

import "fmt"

var nome string

func main() {
	fmt.Println("==== DESAFIO 01 ====")
	fmt.Println("Qual é seu nome?")
	fmt.Scanln(&nome)
	fmt.Printf("Olá %v! prazer em conhecer!", nome)
}