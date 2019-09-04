package main

import (
	"fmt"
)

func Imprimirresultado(nota float64) {
	if nota >= 7  {
		fmt.Println("Aprovado com nota ", nota)
	} else {
		fmt.Println("Reprovado com nota ", nota)
	}
}

func main()  {
	Imprimirresultado(7.3)
	Imprimirresultado(5.1)
}