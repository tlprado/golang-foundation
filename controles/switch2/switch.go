package main

import (
	"time"
	"fmt"
) 

func main() {
	
	t := time.Now()
	
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Bom tarde!")
	default:
		fmt.Println("Boa noite.")
	}
}