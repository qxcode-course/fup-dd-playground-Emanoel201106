package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Print("[ ")
    for ; ; a++ {
        if a % 2 == 0 {
            continue
        }
        if a > b{
            break
        }
        fmt.Printf("%d ", a)

    }
    fmt.Println("]")
}
