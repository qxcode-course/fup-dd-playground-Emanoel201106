package main
import "fmt"
func main() {
    var c, m, soma int
    fmt.Scan(&c)
    for{
        fmt.Scan(&m)
        soma += m
        if soma == 0 {
            fmt.Println("vazio")
        } else if soma < c {
            fmt.Println("ainda cabe")
        } else if soma >= c  && soma < c * 2{
            fmt.Println("lotado")
        } else if soma >= c * 2{
            fmt.Println("hora de partir")
            break
        } 
    }
}
