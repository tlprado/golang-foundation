package main

import "fmt"

func showNumbers(n int) {

	if n > 0 {
		showNumbers(n - 1)
		fmt.Printf("%d ", n)
	}
}

func main() {
	showNumbers(10)
}
