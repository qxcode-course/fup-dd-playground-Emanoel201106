package main
import "fmt"
func main() {
    var c, a int
    fmt.Scan(&c, &a)
    viagens := a / (c - 1)
    if a % (c-1) != 0 {
        fmt.Println(viagens + 1)
    } else{
        fmt.Println(viagens)
    }
}
