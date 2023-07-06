package main

import "fmt"

func main() {
	fmt.Print("Exibe o resultado na mesma linha ")
	fmt.Println("Após exibir o resultado, pula uma linha")
	//retorna o valor 3.25 para a variavel xs
	xs := fmt.Sprint(3.25)
	fmt.Print(xs)
	//Além de exibir o valor da variavel, permite a concatenação
	fmt.Println("O valor de xs é:", xs)
}
