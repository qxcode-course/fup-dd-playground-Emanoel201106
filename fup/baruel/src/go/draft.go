package main
import "fmt"
func main() {
    var qtdtotal, qtdpossui int
    fmt.Scan(&qtdtotal, &qtdpossui)
    var figurinhas []int = make([]int, qtdpossui)

    for i := range figurinhas{
        fmt.Scan(&figurinhas[i])
    }

    repetidas := []int{}

    for i := 0; i < len(figurinhas); i++ {
        for a := 0; a < i; a++ {
            if figurinhas[i] == figurinhas[a] {
                repetidas = append(repetidas, figurinhas[i])
                break
            }
        }
    }

    faltando := []int{}

    for num := 1; num <= qtdtotal; num++ {
        encontrou := false

        for i := 0; i < len(figurinhas); i++ {
            if figurinhas[i] == num {
                encontrou = true
                break
            }
        }

        if !encontrou {
            faltando = append(faltando, num)
        }
    }

    fmt.Print("[ ")

    for _, num := range repetidas {
        fmt.Print(num, " ")
    }

    fmt.Println("]")

    fmt.Print("[ ")

    for _, num := range faltando {
        fmt.Print(num, " ")
    }

    fmt.Println("]")

}