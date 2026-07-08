package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    animais := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&animais[i])
    }

    formado := make([]bool, n)
    casais := 0

    for i := 0; i < n; i++ {
        if formado[i] {
            continue
        }

        for a := i + 1; a < n; a++ {
            if !formado[a] && animais[a] == -animais[i] {
                casais++
                formado[i] = true
                formado[a] = true
                break
            }
        }
    }

    fmt.Println(casais)
}
