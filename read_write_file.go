package main

import (
    "fmt"
    "os"
    "log"
	"bufio"
	"io/ioutil"
)

func main() {
	file, err := os.Open("teste.txt")
	var cpfByte []byte
	var cpfs string

    if err != nil {
        log.Fatal(err)
	}
	
    defer file.Close()

	scanner := bufio.NewScanner(file)

    for scanner.Scan() {

		cpf := scanner.Text()

		lenCpf := len(cpf)

		if(lenCpf < 11) {

			if(lenCpf == 10) {
				cpf = "0" + cpf
				
			} else if(lenCpf == 9) {
				cpf = "00" + cpf

			} else if(lenCpf == 8) {
				cpf = "000" + cpf
			}
		}

		cpfs = cpfs + cpf + "\n"

		fmt.Println(cpf)		
	}

	cpfByte =[]byte(cpfs)

	err = ioutil.WriteFile("cpfsClientesEasynvest.txt", cpfByte, 0777)

	if err != nil {
		log.Fatal(err)
	}
}