package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) toString() {
	fmt.Println(p.nome, p.sobrenome, p.idade)
}

func main() {
	p := pessoa{"José", "Carlos", 34}
	p.toString()
}
