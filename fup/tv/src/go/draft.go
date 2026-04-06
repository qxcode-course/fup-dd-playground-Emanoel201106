package main
import "fmt"
func calcular_valor(valor, parcelas float64) float64{
    valor *= (1+0.05*(parcelas-1))
    return valor
}
func main() {
    var valor float64
    var parcelas float64
    fmt.Scan(&valor, &parcelas)
    valor = calcular_valor(valor, parcelas) 
    cada := valor / parcelas
    fmt.Printf("%.2f\n", cada)
    fmt.Printf("%.2f\n", valor)
}
