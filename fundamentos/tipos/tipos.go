package main

import (
	"math"
    "fmt"
    "reflect"
)

func main() {
	
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Litaral inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64 (uint é um inteiro sem sinal, ou seja, somente os positivos
	
	//byte alias(apelido) para o uint8
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("o Valor maximo do int é", i1)
	fmt.Println("o tipo de i1 é", reflect.TypeOf(i1))

	//numeros reais (float32, float 64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é ", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)); //float64

	//boolean
	bo := true
	fmt.Println("O tipo da variável bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//String
	s1 := "Olá meu nome é Thiago"
	fmt.Println(s1 + "!!!")
	fmt.Println("O tamanho do string é", len(s1))

	//String multiplas linhas
	s2 := `Muitas 
	linhas
	para 
	um grande 
	texto`

	fmt.Println("O tamanho da string é", len(s2))

	//Char
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}