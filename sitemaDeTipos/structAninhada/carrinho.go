package main

import "fmt"

type produto struct {
	IDProduto  int
	titulo     string
	valor      float64
	quantidade int
}

type pedido struct {
	IDPedido int
	itens    []produto
}

// Calculando o valor total do pedido
func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += float64(item.quantidade) * item.valor

	}
	return total
}

func main() {
	pedido1 := pedido{
		IDPedido: 1,
		itens: []produto{
			produto{IDProduto: 1, titulo: "Celular Samsung", valor: 2500.00, quantidade: 2},
		},
	}
	pedido2 := pedido{
		IDPedido: 2,
		itens: []produto{
			produto{IDProduto: 1, titulo: "Celular Samsung", valor: 2500.00, quantidade: 3},
		},
	}

	fmt.Printf("valor total do pedido é R$ %.2f ", pedido1.valorTotal())
	fmt.Println(pedido1)
	fmt.Printf("valor total do pedido é R$ %.2f ", pedido2.valorTotal())
	fmt.Println(pedido2)

}
