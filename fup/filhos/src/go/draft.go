package main
import "fmt"
func main() {
    var idade, qtd int
    fmt.Scan(&idade, &qtd)
    for i := 0; i < qtd; i++ {
        fmt.Println(idade)
        idade += 2
    }
}
