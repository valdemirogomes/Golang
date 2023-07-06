package main

import "fmt"

type imprimivel interface {
	toString() string
}
type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(i imprimivel) {
	fmt.Println(i.toString())
}
func main() {
	var p imprimivel = produto{"Celular Xiomi", 3000}
	imprimir(p)

	p = produto{"Celular Samsung", 3000}
	fmt.Println(p.toString())

	p = pessoa{"Maria", "Cecília"}
	imprimir(p)
	fmt.Println(p.toString())

}
