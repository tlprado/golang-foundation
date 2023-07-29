package main

import "fmt"

func factorial(n int) int {

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func showNumbers(n int) {

	if n > 0 {
		showNumbers(n - 1)
		fmt.Printf("%d ", n)
	}
}

func main() {
	fmt.Println(factorial(3))
}
