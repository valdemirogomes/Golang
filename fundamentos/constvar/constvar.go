package main

import (
	"fmt"
	"math"
)

var raio = 3.2

const PI = 3.1415

func main() {
	area := PI * math.Pow(raio, 2)

	fmt.Println(area)

	const (
		a = 1
		b = 2
	)

	fmt.Println("Constantes", a, b)

	var (
		c = 1
		d = 2
	)
	fmt.Println("Variaveis", c, d)
	var falso, verdadeiro bool = false, true
	fmt.Println(falso, verdadeiro)

	num, bol, str := 353.4, false, "string"

	fmt.Println(num, bol, str)
}
