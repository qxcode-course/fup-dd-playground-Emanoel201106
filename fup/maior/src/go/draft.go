package main
import "fmt"
func main() {
    var chute float64
    var escolha string
    var valor float64
    fmt.Scan(&chute, &escolha, &valor)

    if chute == valor {
        fmt.Println("primeiro")
    } else if escolha == "M" {
        if valor > chute {
            fmt.Println("segundo")
        } else {
            fmt.Println("primeiro")
        }
    } else {
        if valor > chute {
            fmt.Println("primeiro")
        } else {
            fmt.Println("segundo")
        }
    }
}