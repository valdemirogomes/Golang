package main

import "fmt"

func trocar(p2, p1 int) (segundo, primeiro int) {

	segundo = p2
	primeiro = p1
	return
}

func soma(valor1, valor2 float64) (nota1, nota2 float64) {
	nota1 = valor1
	nota2 = valor2

	return
}

func main() {
	x, y := trocar(5, 9)
	fmt.Println(x + y)

	valor1, valor2 := soma(4, 6)

	soma := valor1 + valor2
	fmt.Println(soma)

}
