package main

import (
	"fmt"
)

func main()  {
	var b byte = 3
	fmt.Println(b)

	i := 3 // Inferencia de Tipo (crai a variavél e ja atribui um valor)
	i += 3 // Atribução Aditiva i = i + 2
	i -= 3 // Atribuição Subtrativa i = i - 2
	i /= 2 // Atribuição Divitiva i = i / 2
	i *= 2 // Atribuição Multiplicativa i = i * 2
	i %= 2 //Atribuição de Modulo i = i % 2

	fmt.Println(i)

	x, y := 1, 2

	fmt.Println(x, y)

	x, y = y, x

	fmt.Println(x, y)
}