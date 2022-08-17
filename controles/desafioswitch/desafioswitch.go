package main

import (
	"fmt"
)

func NotaParaConceito(n float64) string {
	switch {
	case  n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	case n >= 0 && n < 3:
		return "E"
	default:
		return "Nota Invalida"
	}
}

func main() {
	fmt.Println(NotaParaConceito(9.8))
	fmt.Println(NotaParaConceito(6.9))
	fmt.Println(NotaParaConceito(2.1))
	fmt.Println(NotaParaConceito(11))
	fmt.Println(NotaParaConceito(4))
}