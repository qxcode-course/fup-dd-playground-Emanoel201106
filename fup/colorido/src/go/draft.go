package main
import "fmt"
func main() {
    var pedra int
    var pe string
    fmt.Scan(&pedra, &pe)
    fmt.Print("[ ")
    passos := 0
    for i := 0; i <= 10; i++ {
        if i == pedra {
            continue
        }
        if i == 10{
            fmt.Print("ceu ")
            continue
        }
        if pe == "d"{
            if passos % 2 == 0 {
                fmt.Printf("%dd ", i)
            } else {
                fmt.Printf("%de ", i)
            }
        } else if pe == "e" {
            if passos % 2 == 0 {
                fmt.Printf("%de ", i)
            } else{
                fmt.Printf("%dd ", i)
            }
        }
        passos++
    }
    fmt.Println("]")
}
