package main

import (
	"math/rand"
	"time"
	"fmt"
)

func NumeroAleatorio() int {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func main()  {
	
	if i := NumeroAleatorio(); i > 5 {
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu!")
	}
}