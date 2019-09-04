package main

import (
	"strconv"
	"fmt"
)

func main()  {
	
	x := 2.4
	y := 2

	fmt.Println( x / float64(y))

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)

	//cuidado ao concatenar
	fmt.Println("Teste código tablea 97 na tabela ASC é: " + string(97))

	//Convertendo int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	//Convertendo string para int
	num, _ := strconv.Atoi("123")

	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") //OU b, _ := strconv.ParseBool("1")

	if(b) {
		fmt.Println("Verdadeiro")
	}
}

