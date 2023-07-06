package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo      string
	turboLigado bool
	velocidade  int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	veiculo01 := ferrari{"F40", false, 234}
	veiculo01.ligarTurbo()
	fmt.Println(veiculo01)

	var veiculo02 esportivo = &ferrari{"F80", false, 236}
	veiculo02.ligarTurbo()
	fmt.Println(veiculo02)
}
