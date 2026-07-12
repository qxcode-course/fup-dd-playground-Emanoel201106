package main
import "fmt"
func main() {
    matriz := make([][]int, 2)

    for i := 0; i < 2; i++ {
        matriz[i] = make([]int, 3)
        for e := 0; e < 3; e++ {
            fmt.Scan(&matriz[i][e])
        }
    }

    soma := 0

    for i := 0; i < 2; i++ {
        for e := 0; e < 3; e++ {
            soma += matriz[i][e]
        }
    }

    fmt.Println(soma)
}