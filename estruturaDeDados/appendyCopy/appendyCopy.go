package main

import "fmt"

func main() {
	//array := [3]int{1, 2, 3}
	var slice []int
	//Acrescentando elementos ao slice usando o append
	slice = append(slice, 4, 5, 6, 7, 8, 9)
	fmt.Println(slice)

	//A makefunção aloca uma matriz zerada e retorna uma fatia que se refere a essa matriz:
	slice2 := make([]int, 3)
	copy(slice2, slice)
	fmt.Println(slice2)
}
