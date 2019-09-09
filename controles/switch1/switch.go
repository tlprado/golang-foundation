package main	

import (
	"fmt"
)

func NotaParaConceito(n float64) string { 
	var nota = int(n)

	switch nota { 
	case 10:
		fallthrough //Executa o case e o dexaixo na seguencia
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota Inválida"
	}
}

func main() {
	fmt.Println(NotaParaConceito(9.8))
	fmt.Println(NotaParaConceito(6.9))
	fmt.Println(NotaParaConceito(2.1))
	fmt.Println(NotaParaConceito(11))
	fmt.Println(NotaParaConceito(4))
}