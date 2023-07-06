package main

import "fmt"

var total float64

func main() {
	media(6.7, 8.9, 7.1, 9.5)
	fmt.Println(media(6.7, 8.9, 7.1, 9.5))
	result := numeros(3, 5, 6, 4, 3, 5, 6, 7, 6, 4, 4, 6, 6, 4)
	fmt.Print("Resultado da soma:", result)

}

/*
Uma função recebe essa nomenclatura quando pode receber um número indefinido de parâmetros.
Para se declarar uma função desse tipo, usamos …
seguido do tipo que esperamos.
Exemplo:
*/

func media(valor ...float64) float64 {
	for _, valores := range valor {
		total += valores

	}
	return total / float64(len(valor))
}

func numeros(num ...int) int {
	fmt.Println("num", num)
	total := 0
	for _, numeros := range num {
		//Exibi os valores
		fmt.Println(numeros)

		// Realiza a soma dos valores passado na função
		total += numeros
	}
	return total

}
