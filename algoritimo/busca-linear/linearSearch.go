package main

import "fmt"

func linearSearch(numbers [9]int, n int) string {

	for i := 0; i < len(numbers); i++ {
		if numbers[i] == n {
			return "True"
		}
	}

	return "False"
}

func main() {

	array := [9]int{1, 8, 32, 91, 5, 15, 9, 100, 3}
	fmt.Print(linearSearch(array, 91))
}
