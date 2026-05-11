package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Print("[ ")
    for i := a; i <= b; i++ {
        for i := b; i > a; i-- {
            fmt.Printf("%d ", i)
        }
        fmt.Printf("%d ",i)
    }
    fmt.Println("]")
}
