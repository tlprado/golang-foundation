package main

import (
	"fmt"
)

func main()  {
	i := 1

	var p *int = nil

	p = &i //pegando o endereço da variavél e atribuindo para o ponteiro
	// &i Endereço de memória 

	*p++ //desrefenciando o ponteiro para pegar o valor
	i++

	// p++  Go não tem aritimética de ponteiros

	fmt.Println(p, *p, i, &i)
}