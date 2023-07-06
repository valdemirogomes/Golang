package main

import "fmt"

func main() {
	x := 1
	// Definindo a variavel y como um ponteiro
	var y *int
	// Apontando a variavel y para o endereço de memoria da variavél x
	y = &x
	// Imprimindo o valor da variavel y
	fmt.Println(*y)
	// Imprimindo o endereço de memoria da variável x e y
	fmt.Println(&x, &y)

}
