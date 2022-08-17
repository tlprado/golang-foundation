package main

import (
	"fmt"
	"reflect"
)

func main() {
	array1 := [3]int{1, 2, 3} //array
	slice1 := []int{1, 2, 3}  //slice não é um array, mas sim uma fatida/pedaço de um array.

	fmt.Println(array1, slice1)

	fmt.Println(reflect.TypeOf(array1), reflect.TypeOf(slice1))

	array2 := [5]int{1, 2, 3, 4, 5}

	slice2 := array2[1:3]
	fmt.Println(array2, slice2)

	slice3 := array2[:2]
	fmt.Println(array2, slice3)

	slice4 := slice2[:1]
	fmt.Println(slice2, slice4)

}
