package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[231234] = "José"
	aprovados[908678] = "Maria"
	aprovados[332123] = "Niclaus"

	//Percorrendo com for range
	for cpf, nome := range aprovados {
		fmt.Println(cpf, nome)

	}
	//Percorrendo com for tradicional
	result := ""
	for i := 0; i < len(aprovados); i++ {
		result += aprovados[i]
	}
	fmt.Println(result)

}
