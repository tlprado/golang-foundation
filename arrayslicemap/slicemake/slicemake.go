package main

import "fmt"

func main() {
	slice := make([]int, 10)
	slice[9] = 12
	fmt.Println(slice)

	slice = make([]int, 10, 20)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 1)
	fmt.Println(slice, len(slice), cap(slice))
}
