package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    paredes := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&paredes[i])
    }

    cont := 0
    maior := 0

    for i := 0; i < n; i++{
        if paredes[i] > maior {
            cont++
            maior = paredes[i]
        }
    }

    fmt.Println(cont)
}