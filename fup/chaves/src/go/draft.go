package main
import "fmt"
func main() {
    var num int
    fmt.Scan(&num)

    if num > 0 {
        fmt.Println("positivo")
    } else if num < 0 {
        fmt.Println("negativo")
    } else {
        fmt.Println("nulo")
    }
}