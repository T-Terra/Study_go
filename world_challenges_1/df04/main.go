package main

import (
	"example/world/world_challenges_1/df04/parser"
	"fmt"
)

func main() {
	fmt.Println("==== DESAFIO 04 ====")
	var char rune
	fmt.Print("Digite um caractere: ")
	fmt.Scanf("%c", &char)
	parser.Convert(char)
}
