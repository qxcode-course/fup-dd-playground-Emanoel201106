package main
import "fmt"
func main() {
    var qtd int
    fmt.Scan(&qtd)
    
    chocolate := 0
    limao := 0
    manha := 0
    tarde := 0

    for i := 0; i < qtd; i++{
    var sabor, turno string
    fmt.Scan(&sabor, &turno)

    if sabor == "c" {
        chocolate++
    } else {
        limao++
    }
    if turno == "m" {
        manha++
    } else {
        tarde++
    }
    }
    if chocolate > limao {
        fmt.Println("c")
    } else if limao > chocolate {
        fmt.Println("l")
    } else {
        fmt.Println("empate")
    }
    if manha < tarde {
        fmt.Println("m")
    } else if tarde < manha {
        fmt.Println("t")
    } else {
        fmt.Println("empate")
    }
    
}