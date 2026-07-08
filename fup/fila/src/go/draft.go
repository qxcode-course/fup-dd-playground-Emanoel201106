package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    alunos := []int{}
    servidores := []int{}

    for i := 0; i < n; i++ {
        var x int
        fmt.Scan(&x)

        if x % 2 == 0 {
            servidores = append(servidores, x)
        } else {
            alunos = append(alunos, x)
        }
    }
    fmt.Print("[ ")
    for i := 0; i < len(alunos); i++ {
        fmt.Print(alunos[i], " ")
    }
    fmt.Println("]")

    fmt.Print("[ ")
    for i := 0; i < len(servidores); i++ {
        fmt.Print(servidores[i], " ")
    }
    fmt.Println("]")
}