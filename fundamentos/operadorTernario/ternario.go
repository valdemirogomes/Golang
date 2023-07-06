package main

import "fmt"

// Não existe operador ternário na linguagem go
func media(nota int) string {
	if nota <= 6 {
		return "Reprovado"
	}
	return "Aprovado"

}

func main() {
	resultado := media(7)
	fmt.Println(resultado)
}
