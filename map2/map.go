package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jão Zé": 11000.80,
		"Mary":   15600.55,
		"Piter":  2500.0,
	}

	funcsESalarios["Rafa"] = 2800.90

	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
