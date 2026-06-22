package main
import "fmt"
func main() {
    var m int
    fmt.Scan(&m)

    anterior := 0
    fmt.Scan(&anterior)
    var x int
    parkour := 0
    for i := 1; i < m; i++ {
        fmt.Scan(&x)
        if x - anterior > 1 || x - anterior < -1 {
            anterior = x
            parkour++
        } else {
            anterior = x
            continue
        }
    }
    fmt.Println(parkour)
}