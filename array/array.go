package main

import "fmt"

func main() {
	var notas [3]float64
	notas[0], notas[1], notas[2] = 34.56, 56.7, 87.6
	x := 0.0
	for i := 0; i < len(notas); i++ {
		x = notas[i]
		fmt.Println("Posição", i, "=", x)

	}
	media()
}

func media() {
	var pontos [4]float64

	pontos[0], pontos[1], pontos[2], pontos[3] = 8.8, 8.0, 9.1, 9.0
	contador := 0.0
	mediaFinal := 0.0
	for i := 0; i < 4; i++ {
		contador += pontos[i]
		mediaFinal = contador / 4

	}
	fmt.Println(mediaFinal)
}
