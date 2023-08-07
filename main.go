package main

import "fmt"

type meutipo int
var x meutipo
var y int

func main() {
    fmt.Printf("Valor: %v\n", x)
    fmt.Printf("Tipo: %T\n", x)
    x = 42
    fmt.Printf("Valor: %v\n", x)
    y = int(x)
    fmt.Printf("Valor: %v\n", y)
    fmt.Printf("Tipo: %T\n", y)
}