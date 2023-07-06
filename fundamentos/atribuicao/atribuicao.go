package main

import "fmt"

func main() {
	// definindo a variavel x do tipo int e inicializando com valor 10
	var x int = 10
	fmt.Println(x)

	a := 3 // inferencia de tipo
	fmt.Println(a)

	a += 1 // a = a + 1
	fmt.Println(a)

	a -= 1
	fmt.Println(a)
	//:= definindo e inicializando a variavel b com valor 10 e c com valor 20
	b, c := 10, 20
	fmt.Println(b, c)
	// = atribuindo um novo valor as variaveis b e c
	b, c = c, b
	fmt.Println(b, c)

}
