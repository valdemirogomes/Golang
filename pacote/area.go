package main

type Retangulo struct {
	base   int
	altura int
}

func area(r Retangulo) (result int) {
	result = r.base * r.altura
	return result
}
