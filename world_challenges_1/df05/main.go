package main

import "fmt"

func main() {
	var n int
	var a int
	var s int
	
	fmt.Println("==== DESAFIO 05 ====")
	fmt.Print("Digite um número: ")
	fmt.Scanln(&n)
	a = n - 1
	s = n + 1
	fmt.Printf("O antecessor de %v, é %v e o sucessor é %v", n, a, s)
}