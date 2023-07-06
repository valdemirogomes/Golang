package main

import "fmt"

type carro struct {
	nome       string
	velocidade string
}
type ferrari struct {
	carro // campos anonimos
	turbo bool
}

func (f ferrari) mostrarVeiculo() {
	fmt.Println("Veiculo:", f.nome, "Velocidade:", f.velocidade, "turbo:", f.turbo)

}

func main() {
	ferrari := ferrari{}
	ferrari.nome = "F40"
	ferrari.velocidade = "320 Km/h"
	ferrari.turbo = true
	ferrari.mostrarVeiculo()
}
