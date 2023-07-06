package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	RG        int
}

type professor struct {
	pessoa
	matricula int
	formacao  string
}

type aluno struct {
	pessoa
	serie string
}

func main() {
	professor := professor{}
	professor.nome = "Alfredo"
	professor.sobrenome = "Silva"
	professor.RG = 3445678
	professor.matricula = 3443209
	professor.formacao = "Bacharel em Letras - Portugues"
	fmt.Println(professor)

	aluno := aluno{
		// aluno é uma pessoa
		pessoa: pessoa{
			nome:      "José",
			sobrenome: "Cabral",
			RG:        3443907,
		},
		serie: "Segundo ano do ensino médio",
	}
	fmt.Println(aluno)

}
