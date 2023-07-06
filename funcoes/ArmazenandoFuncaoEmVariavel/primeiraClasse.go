package main

import "fmt"

func main() {
	result := soma(20, 7)

	sub(10, 3)

	fmt.Println(soma(20, 7))
	fmt.Println(result)
	fmt.Println(media(4.1, 3.0, 4.4, 5.5))

}

//armazenando funcao anonima em variavel

var soma = func(x, y int) int {

	return x + y
}

var sub = func(x, y int) {
	resultSub := x - y

	fmt.Println(resultSub)
}

//Função como expressão

var media = func(nota ...float64) float64 {
	total := 0.0
	for _, notas := range nota {
		total += notas

	}

	return total / float64(len(nota))

}
