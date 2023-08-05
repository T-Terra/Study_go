package main

import "fmt"

var dia int
var mes string
var ano int

func main() {
	fmt.Println("==== DESAFIO 02 ====")
	fmt.Println("Dia:")
	fmt.Scanln(&dia)
	fmt.Println("Mês:")
	fmt.Scanln(&mes)
	fmt.Println("Ano:")
	fmt.Scanln(&ano)

	fmt.Printf("Você nasceu do dia %v de %v de %v correto?\n", dia, mes, ano)
}