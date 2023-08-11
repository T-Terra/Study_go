package main

import (
	"fmt"
	"math"
)
var v float64

func main(){
	fmt.Println("==== DESAFIO 06 ====")
	fmt.Print("Digite um número: ")
	fmt.Scanln(&v)
	d := v * 2
	t := v * 3
	raiz := math.Sqrt(v)

	fmt.Printf("O dobro de %v é %v, o triplo é %v e a raiz quadrada é %v", v, d, t, raiz)

}