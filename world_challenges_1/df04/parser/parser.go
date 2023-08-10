package parser

import (
	"fmt"
	"unicode"
	"time"
)

func Convert(char rune) {
	if unicode.IsDigit(char) {
		fmt.Println("O caractere é um dígito numérico.")
	} 
	if unicode.IsLetter(char) {
		fmt.Println("O caractere é uma letra.")
	} 
	if unicode.IsLower(char) {
		fmt.Println("O caractere é minúsculo.")
	}
	if unicode.IsUpper(char) {
		fmt.Println("O caractere é maiúsculo.")
	}
	if unicode.IsSpace(char) {
		fmt.Println("O caractere é um espaço")
	}

	time.Sleep(10 * time.Second)
}

