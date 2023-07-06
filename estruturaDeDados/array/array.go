package main

import "fmt"

func main() {

	var x [5]int
	x[0], x[1], x[2], x[3], x[4] = 1, 2, 3, 4, 5
	fmt.Println(x)
	tamanho := len(x)
	fmt.Println(tamanho)
}
