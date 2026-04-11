package main
import "fmt"
func main() {
    var dia, hora int
    fmt.Scan(&dia, &hora)
    if dia > 1 && hora >= 8 && hora <=11 && dia < 7 || dia > 1 && hora >= 14 && hora <=17 && dia <7 {
        fmt.Println("SIM")
    } else if dia == 7 && hora >= 8 && hora <= 11 {
        fmt.Println("SIM")
    } else {
        fmt.Println("NAO")
    }
}
