package main

import "fmt"

func nota(valor float64) {
	if valor <= 6 {
		fmt.Println("Reprovado")
	} else {
		fmt.Println("Aprovado")
	}
}

func main() {
	nota(6.1)
}
