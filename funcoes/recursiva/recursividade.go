package main

import "fmt"

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
func main() {
	fmt.Println(fatorial(1))
	fmt.Println(fibonacci(15))

}
