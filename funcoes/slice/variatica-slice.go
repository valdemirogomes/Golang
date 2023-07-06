package main

import "fmt"

func main() {
	aprovado := []string{"Maria", "Yasmin", "Heloisa", "Ketlen"}
	aprovados(aprovado...)

}

func aprovados(alunos ...string) {
	fmt.Println("Aprovados")

	for i, aluno := range alunos {
		fmt.Println(i+1, aluno)

	}
}
