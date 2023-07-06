package main

import "fmt"

func main() {

	fmt.Println(media(6))
	fmt.Println(horaDia(19))

}

func media(valor float64) string {

	switch valor {
	case 4, 3, 2, 1:
		return "Reprovado"
	case 8, 9, 10:
		return "Aprovado"
	case 5, 6, 7:
		return "Recuperação"
	default:
		return "Nota inválida"

	}

}

func horaDia(valor int) string {
	switch {
	case valor <= 12 && valor >= 1:
		return "Bom dia"
	case valor >= 12 && valor <= 18:
		return "Boa tarde"
	default:
		return "Boa noite"
	}
}
