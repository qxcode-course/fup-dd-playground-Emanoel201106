package main
import "fmt"
func main() {
    var nota1, nota2, af int
    fmt.Scan(&nota1, &nota2, &af)
    media := (nota1 + nota2) / 2
    if media >= 7 {
        fmt.Println("aprovado")
    } else if media < 4 {
        fmt.Println("reprovado")
    } else if media >= 4 && media < 7 {
        media2 := (media + af) / 2
        if media2 >= 5 {
            fmt.Println("aprovado na final")
        } else {
            fmt.Println("reprovado na final")
        }
    }
}
