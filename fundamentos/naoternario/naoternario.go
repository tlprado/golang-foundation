package main

import (
	"fmt"
) 

// Não tem operador tenário em Go
func ObterResultado(nota float64) string  {
	if(nota >= 6) {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(ObterResultado(6.2))
}