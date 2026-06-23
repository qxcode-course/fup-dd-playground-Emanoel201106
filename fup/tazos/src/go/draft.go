package main
import "fmt"
func main() {
    var qtd int
    fmt.Scan(&qtd)
    
    colecao := make([]int, qtd)
    for i := range colecao {
        fmt.Scan(&colecao[i])
    }
    
}