package main
import "fmt"
func main() {
    var comp1, lar1, comp2, lar2 int
    fmt.Scan(&comp1, &lar1, &comp2, &lar2)

    area1 := comp1 * lar1
    area2 := comp2 * lar2

    if area1 > area2 {
        fmt.Println(area1)
    } else {
        fmt.Println(area2)
    }
}