package main
import "fmt"
func main() {
    var qtd int
    fmt.Scan(&qtd)
    var n []int = make([]int, qtd)
    for i := range n{
        fmt.Scan(&n[i])
    }
    fmt.Print("[ ")
    for _, valor := range n{
        fmt.Printf("%d ", valor)
    }
    fmt.Println("]")
}
