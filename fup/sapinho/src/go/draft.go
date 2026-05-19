package main
import "fmt"
func main() {
    var p, s, e int
    fmt.Scan(&p, &s, &e)
    posicao := 0
    fmt.Printf("%d", posicao)
    sapo := 0
    for {
        sapo = sapo + s
        if sapo >= p {
            fmt.Println(" saiu")
            break
        }
        fmt.Println("", sapo)
        sapo = sapo - e
        fmt.Printf("%d", sapo)
    }
}
