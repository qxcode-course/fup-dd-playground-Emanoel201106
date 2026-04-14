package main
import "fmt"
func mostrar_vetor(n []int, sep string) {
    for i, valor := range n {
        if i != 0 {
            fmt.Println(sep)
        }
        fmt.Printf("%d", valor)
    }
    fmt.Println("")
}
func main() {
    var qtd int
    fmt.Scan(&qtd)
    
    var n []int = make([]int, qtd)
    for i := range n{
        fmt.Scan(&n[i])
    }
    mostrar_vetor(n, "")
}
