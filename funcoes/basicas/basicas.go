package main

import "fmt"

func main() {
	imprimir()
}
func imprimir() {
	fmt.Println(soma(10, 3))
	a, b := vlrs(30, 50)
	fmt.Println(a, b)

}

func soma(nota1 int, nota2 int) int {
	resultado := nota1 + nota2
	return resultado
}
func vlrs(x, y int) (int, int) {
	return x, y
}
