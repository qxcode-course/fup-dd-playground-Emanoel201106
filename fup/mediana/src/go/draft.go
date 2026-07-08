package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    vetor := make([]float64, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&vetor[i])
    }

    for i := 0; i < n-1; i++ {
        for a := 0; a < n-1-i; a++ {
            if vetor[a] > vetor[a+1] {
                vetor[a], vetor[a+1] = vetor[a+1], vetor[a]
            }
        }
    }

    var mediana float64

    if n % 2 != 0 {
        mediana = vetor[n/2]
    } else {
        mediana = (vetor[n/2-1] + vetor[n/2]) / 2
    }
    fmt.Printf("%.1f\n", mediana)
}