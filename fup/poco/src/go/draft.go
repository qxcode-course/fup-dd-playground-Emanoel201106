package main
import "fmt"
func main() {
    var p, s, e int
    fmt.Scan(&p, &s, &e)

    posicao := 0

    for {
        novo := posicao + s
        
        if novo >= p {
            fmt.Println(posicao, "saiu")
            break
        }
        fmt.Println(posicao, novo)

        posicao = novo - e

        if posicao < 0{
            fmt.Println(posicao, "morreu")
            break
        }

        if s > 0 {
            s -= 10
        }
    }
}