package main
import "fmt"
func main() {
    var n int 
    fmt.Scan(&n)

    matriz := make([][]int, n)

    for i := 0; i < n; i++ {
        matriz[i] = make([]int, n)
        for e := 0; e < n; e++ {
            fmt.Scan(&matriz[i][e])
        }
    }

    maior :=  -1
    indice := 0

    for e := 0; e < n; e++ {
        soma := 0

        for i := 0; i < n; i++ {
            soma += matriz[i][e] * matriz[i][e]
        }

        if soma > maior {
            maior = soma
            indice = e
        }
    }
    fmt.Println(indice)
}