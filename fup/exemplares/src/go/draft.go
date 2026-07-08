package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    animais := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&animais[i])
    }

    unicos := []int{}

    for i := 0; i < n; i++ {
        existe := false

        for a := 0; a < len(unicos); a++ {
            if animais[i] == unicos[a] {
                existe = true
                break
            }
        }
        if !existe {
            unicos = append(unicos, animais[i])
        }
    }
    
    for i := 0; i < len(unicos) - 1; i++ {
        for b := 0; b < len(unicos) - 1 - i; b++ {
            if unicos[b] > unicos[b+1] {
                unicos[b], unicos[b+1] = unicos[b+1], unicos[b]
            }
        }
    }

    for i := 0; i < len(unicos); i++ {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(unicos[i])
    }
    fmt.Println()
}