package main

import "fmt"

func main() {

	pessoa := map[string]float64{

		"João":      3.500,
		"Maria":     4.560,
		"Francisco": 2.500,
	}
	for nome, salario := range pessoa {
		fmt.Println(nome, salario)
	}

}
