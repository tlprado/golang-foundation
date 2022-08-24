package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"A": {
			"Adailzo":  18000.50,
			"Amarildo": 18000.50,
		},
		"B": {
			"Bianca":  25000.50,
			"Brauler": 15000.50,
		},
		"C": {
			"Clara":   30000.50,
			"Carlito": 15000.50,
		},
	}

	delete(funcsPorLetra, "C")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
