package main
import "fmt"
func main() {
    var m, a, b int
    fmt.Scan(&m, &a, &b)
    c := m - (a + b)
    if c > a && c > b {
        fmt.Println(c)
    } else if a > b && a > c {
        fmt.Println(a)
    } else {
        fmt.Println(b)
    }
}
