package main
import "fmt"
func main() {
    vetor := make([]int, 5)

    for i := 0; i < 5; i++ {
        fmt.Scan(&vetor[i])
    }

    maior := vetor[0]
    menor := vetor[0]

    for i := 1; i < 5; i++ {
        if vetor[i] > maior {
            maior = vetor[i]
        }
        if vetor[i] < menor {
            menor = vetor[i]
        }
    }
    fmt.Println(maior + menor)
}