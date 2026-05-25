package main
import "fmt"
func mostrar_vetor(n []int, sep string) {
    for i, valor := range n {
        if i != 0 {
            fmt.Printf("%s, ", sep)
        }
        fmt.Print(valor)
    }
}
func main() {
    var qtd int
    fmt.Scan(&qtd)
    
    var n []int = make([]int, qtd)
    for i := range n{
        fmt.Scan(&n[i])
    }
    fmt.Print("[")
    mostrar_vetor(n, "")
    fmt.Println("]")
}
