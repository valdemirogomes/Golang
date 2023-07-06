package main

import "fmt"

func main() {
	x := 10
	y := 2.5

	z := float64(x) / y

	fmt.Println(z)

	var a float64 = 10
	b := a == float64(x)
	fmt.Println(b)
	type double float64

	var c double = 30.6
	fmt.Println(c)
	d := double(a) == c
	fmt.Println(d)

}
