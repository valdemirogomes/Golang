package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
}

func (p *pessoa) setNome(nome, sobrenome string) {
	p.nome = nome
	p.sobrenome = sobrenome

}

func (p pessoa) getNome() string {
	return p.nome + p.sobrenome
}

func main() {
	pessoa := pessoa{"Maria ", "José"}
	fmt.Println(pessoa.getNome())
	pessoa.setNome("Maria ", "Aparecida")
	fmt.Println(pessoa.getNome())
}
