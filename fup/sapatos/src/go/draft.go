package main
import "fmt"
func main() {
    var a, b, soma int
    fmt.Scan(&a, &b)
    if a > b {
        fmt.Println("invalido")
        return
    }
    for ; a <= b; a++ {
        if a % 2 == 0 {
            soma += a
        }
    }
    fmt.Println(soma)
}
