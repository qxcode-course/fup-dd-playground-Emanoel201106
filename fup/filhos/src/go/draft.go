package main
import "fmt"
func main() {
    var idade, qtd int
    fmt.Scan(&idade, &qtd)
    for ; idade < qtd; idade += 2 {
        
        fmt.Println(idade)
    }
}
