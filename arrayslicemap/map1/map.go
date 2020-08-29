package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[12345666558] = "João"
	aprovados[23665655655] = "Joana"
	aprovados[45665444455] = "Marcia"

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[45665444455])

	delete(aprovados, 45665444455)

	fmt.Println(aprovados[45665444455])
}
