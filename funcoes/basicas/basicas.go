package main

import "fmt"

func PrintWord() {
	fmt.Println("Apple")
}

func PrintFullName(name string, lastName string) {
	fmt.Printf("FullName %s %s\n", name, lastName)
}

func returnText() string {
	return "Orange"
}

func returnFullName(name, lastName string) string {
	return fmt.Sprintf("FullName %s %s\n", name, lastName)
}

func returnTwoResults() (string, string) {
	return "Zezinho", "Donald"
}

func main() {
	PrintWord()
}
