package main

import "fmt"

func main() {
	fmt.Println("Inicializando função 1")
	f1()
}

func f2() {
	fmt.Println("Chamando função 3")
	f3()
}

func f1() {
	fmt.Println("Chamando função 2")
	f2()

}

func f3() {

	fmt.Println("Pilha ")

}
