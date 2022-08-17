package main

import (
	"fmt"
	"math"
)

func main() {

	a := 3
	b := 2

	fmt.Println("Soma => ", a+b)

	fmt.Println("Subtração => ", a-b)

	fmt.Println("Divisão => ", a/b)

	fmt.Println("Multplicação => ", a*b)

	fmt.Println("Módulo => ", a%b)

	// bitwise
	fmt.Println("E => ", a&b) // Valor Binário 10 & 11

	fmt.Println("OU => ", a|b) // Valor Binário 10 | 11

	fmt.Println("Xor => ", a^b) // Valor Binário 10 ^ 11


	c := 3.0
	d := 2.0

	//Outras Operações usando Math...
	fmt.Println("Maior Valor entre =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor Valor entre =>", math.Min(c, d)) 
	fmt.Println("Expodenciação =>", math.Pow(c, d)) 
}