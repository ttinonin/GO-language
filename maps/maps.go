package main

import (
	"fmt"
)

func main() {
	produtos := make(map[string]float64)

	produtos["Celular"] = 899.90
	produtos["PC Gamer"] = 4999.90
	produtos["Mouse"] = 299.90
	produtos["Teclado"] = 89.90

	for key, value := range produtos {
		// Formata o valor
		formatado := fmt.Sprintf("%.2f", (value - (value*0.20)))

		fmt.Println("Produto:", key, "=>", "Preço: R$", value, "- Desconto (20%): ", formatado)
	}
}
