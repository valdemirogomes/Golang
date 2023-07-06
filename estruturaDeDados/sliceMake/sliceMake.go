package main

import "fmt"

func main() {
	//Array interno com 10 elementos
	slice := make([]int, 10)
	slice[9] = 12
	fmt.Println(slice)
	//Slice com 10 posições porém com um array interno de tamanho 20
	slice = make([]int, 10, 20)
	fmt.Println("Tamanho =", len(slice), "Capacidade =", cap(slice))
	//Preenchendo a capacidade maxima de 20 elementos
	slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	//Atribuindo um novo valor ao index 19
	slice[19] = 0
	fmt.Println(slice)
	//Com a capacidade máxima preenchida, o tamanho do slice passa a ser 20
	fmt.Println(len(slice), cap(slice))

	//Apend
	array := [3]int{1, 2, 3}
	var slice1 []int

	slice1 = append(slice1, 1, 2, 3)
	fmt.Println(array)
	fmt.Println(slice1)

	//Copy
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2)
}
