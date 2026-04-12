package main
import "fmt"
func main() {
    var m, a, b int
    fmt.Scan(&m, &a, &b)
    idade := m - (a + b)
    if idade > a && idade > b {
        fmt.Println(idade)
    } else if idade < a && a > b || a == b || idade == a {
        fmt.Println(a)
    } else if idade < b && b > a || idade == b {
        fmt.Println(b)
    }  
}
