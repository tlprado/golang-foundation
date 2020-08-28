package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4} //compilador faz a contar e gerar a quantidade correta de elementos

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros { //fro sem usar o incremento
		println(num)
	}
}
