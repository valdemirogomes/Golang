package main

import "fmt"

func sequencia(x int) func() int {
	i := x
	return func() int {
		i += 1
		return i
	}
}
func main() {
	next := sequencia(4)
	fmt.Println(next())
	fmt.Println(next())

}
