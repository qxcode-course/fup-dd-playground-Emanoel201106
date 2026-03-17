package main

import "fmt"

func main() {
	var produto1, produto2, produto3 float64
	var valor1, valor2, valor3 float64
	var dinheiro float64

	fmt.Scan(&produto1, &produto2, &produto3, &valor1, &valor2, &valor3, &dinheiro)
	valor1 = valor1 * produto1
	valor2 = valor2 * produto2
	valor3 = valor3 * produto3

	fmt.Printf("%.2f\n", dinheiro-(valor1+valor2+valor3))

}
