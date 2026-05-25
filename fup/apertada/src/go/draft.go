package main
import "fmt"
func main() {
    var n []int = make([]int, 5)
    for i := range n {
        fmt.Scan(&n[i])
    }
    menor := n[0]
    for _, valor := range n {
        if valor < menor {
        menor = valor
    }
    }
    
    fmt.Println(menor)
}
