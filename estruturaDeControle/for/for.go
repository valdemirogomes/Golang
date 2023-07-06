package main

import "fmt"

func main() {
	for i := 1; i <= 10; i += 1 {
		if i%2 == 0 {
			fmt.Println(i, "Par")

		} else {
			fmt.Println(i, "Impar")

		}

	}

	contador := 1
	fmt.Println("-----------------------------------------------------")
	for contador <= 10 {
		contador += 2
		if contador%2 == 0 {

			fmt.Println(contador, "Par")
		} else {
			fmt.Println(contador, "Impar")
		}

	}
}
