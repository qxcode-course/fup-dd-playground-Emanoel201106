package main
import "fmt"
import "math"
func main() {
    var a, b, c float64
    fmt.Scan(&a, &b, &c)
    delta := math.Pow(b, 2) - 4 * a * c
    x1 := (-b + math.Sqrt(delta)) / (2 * a)
    x2 := (-b - math.Sqrt(delta)) / (2 * a)
    if delta > 0 {
      fmt.Printf("%.2f\n%.2f\n", x1, x2)  
    } else if delta == 0 {
        fmt.Printf("%.2f\n", x1)
    } else {
        fmt.Println("nao ha raiz real")
    }
}
