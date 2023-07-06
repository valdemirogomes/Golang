package main

import "fmt"

func soma(x int, y int) int {

	return x + y
}

func imprimirSoma(valor int) {
	fmt.Print(valor)
}

func main() {
	resultado := soma(5, 6)
	imprimirSoma(resultado)

}
