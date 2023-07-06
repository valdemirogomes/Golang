package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}         //array
	slice := []int{1, 2, 3, 4, 5, 6} //slice
	fmt.Println(array)
	slice[0] = 20
	fmt.Println(slice[0])
	fmt.Println(len(slice))
}
