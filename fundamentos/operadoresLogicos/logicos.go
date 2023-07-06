package main

import "fmt"

func main() {
	tv50, tv32, sorvete := compras(false, true)

	fmt.Printf("Tv50: %t, Tv32: %t, sorvete: %t", tv50, tv32, sorvete)
}

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2
	tomarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, tomarSorvete

}
