package main

import "fmt"

func nota(valor float64) {
	if valor >= 1 && valor < 5 {
		fmt.Println("Reprovado")

	} else if valor >= 5 && valor < 7 {
		fmt.Println("Recuperação")

	} else if valor >= 7 {
		fmt.Println("Aprovado")

	} else {
		fmt.Println("Nota insuficiente")
	}
}

func main() {

	nota(0.6)

}
