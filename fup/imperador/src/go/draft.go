package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    matriz := make([][]string, n)
    linhaLeao, colLeao := -1, -1

    for i := 0; i < n; i++ {
        matriz[i] = make([]string, n)
        for e := 0; e < n; e++ {
            fmt.Scan(&matriz[i][e])

            if matriz[i][e] == "L" {
                linhaLeao = i
                colLeao = e
            }
        }
    }
    g := 0
    c := 0

    for i := 0; i < n; i++ {
        for e := 0; e < n; e++ {
            if i == linhaLeao || e ==colLeao {
                continue
            }

            if matriz[i][e] == "G" {
                g += 2
            } else if matriz[i][e] == "C" {
                if i + e == n-1 {
                    c += 2
                } else {
                    c += 1
                }
            }
        }
    }
    if g > c {
        fmt.Println("Gladiadores")
    } else if c > g {
        fmt.Println("Condenados a morte")
    } else {
        fmt.Println("Ninguem")
    }
}