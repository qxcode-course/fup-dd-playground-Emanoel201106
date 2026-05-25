package main
import "fmt"
func main() {
    var chico, cebolinha, qtd int
    fmt.Scan(&chico, &cebolinha, &qtd)
    animais := make([]string, qtd)
    patas := 0

    for i := range animais {
        fmt.Scan(&animais[i])
        if animais[i] == "v" {
            patas += 4
        } else if animais[i] == "g" {
            patas += 2
        } else if animais[i] == "c" {
            patas += 4
        }
    }
    fmt.Println(patas)

    distChico := chico - patas
    if distChico < 0 {
        distChico = -distChico
    }
    distCebolinha := cebolinha - patas
    if distCebolinha < 0 {
        distCebolinha = -distCebolinha
    }
    
    if distChico < distCebolinha {
        fmt.Println("Chico Bento")
    } else if distCebolinha < distChico {
        fmt.Println("Cebolinha")
    } else {
        fmt.Println("empate")
    }
}
