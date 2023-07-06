package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := produto{1, "Televisao", 6.700, []string{"promocao"}}
	json, _ := json.Marshal(p1)
	fmt.Println(string(json))
}
