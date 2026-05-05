package main
import "fmt"
func main() {
    var pedra int
    fmt.Print("[ ")
    fmt.Scan(&pedra)
    ceu := 10
    for i := 0; i <= ceu; i++ {
        if i == pedra{
            continue
        } 
        if i == ceu{
            
        }
        fmt.Printf("%d ", i)
    }
    fmt.Println("]")
}
