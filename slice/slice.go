package main

import "fmt"

func main() {
	array := [4]int{1, 2, 3, 4}
	slice := []int{1, 2, 3, 4, 5}

	/*
		inicia na posicao de indice 0 ate a posicao de indice 3, sem incluir o 3
		1,2,3,4
		0,1,2
	*/
	slice2 := array[:3]
	fmt.Println(slice, slice2)
	/*
		//inicia na posicao de indice 0 ate a posicao de indice 2, sem incluir o 2
		1,2,3,4
		0,1
	*/
	slice3 := array[:2]
	fmt.Println(slice3)
}
