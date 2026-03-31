package main
import "fmt"
func main() {
    var P, D1, D2 int
    
    fmt.Scan(&P, &D1, &D2)
    soma := D1 + D2
    if soma%2 == P{
        fmt.Println("0")
    } else{
        fmt.Println("1")
    }
}
