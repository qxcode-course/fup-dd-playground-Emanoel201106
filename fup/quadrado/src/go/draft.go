package main
import "fmt"
func main() {
    matriz := make([][]int, 3)

    for i := 0; i < 3; i++ {
        matriz[i] = make([]int, 3)
        for e := 0; e < 3; e++ {
            fmt.Scan(&matriz[i][e])
        }
    }

    ref := matriz[0][0] + matriz[0][1] + matriz[0][2]

    for i := 0; i < 3; i++ {
        soma := 0
        for e := 0; e < 3; e++ {
            soma += matriz[i][e]
        }

        if soma != ref {
            fmt.Println("nao")
            return
        }
    }

    for e := 0; e < 3; e++ {
        soma := 0
        for i := 0; i < 3; i++ {
            soma += matriz[i][e]
        }

        if soma != ref {
            fmt.Println("nao")
            return
        }
    }

    soma := 0
    for i := 0; i < 3; i++ {
        soma += matriz[i][i]
    }

    if soma != ref {
        fmt.Println("nao")
        return
    }

    soma = 0
    for i := 0; i < 3; i++ {
        soma += matriz[i][2-i]
    }

    if soma != ref {
        fmt.Println("nao")
        return
    }

    fmt.Println("sim")
}