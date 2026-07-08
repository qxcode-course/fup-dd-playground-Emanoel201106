package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    cartas := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&cartas[i])
    }

    fmt.Print("[")
    
    for i := 0; i < n; i++ {
        if i > 0 {
            fmt.Print(", ")
        }

    if cartas[i] == 1 {
        fmt.Print("A")
    } else if cartas[i] == 11 {
        fmt.Print("J")
    } else if cartas[i] == 12 {
        fmt.Print("Q")
    } else if cartas[i] == 13 {
        fmt.Print("K")
    } else {
        fmt.Print(cartas[i])
    }

    
    }
    fmt.Println("]")
}