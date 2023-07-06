package main

import "fmt"

func main() {
	soma := 0
	numeros := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, numero := range numeros {
		soma += numero
		fmt.Println(i, numero)

	}
	fmt.Println(soma)
}
