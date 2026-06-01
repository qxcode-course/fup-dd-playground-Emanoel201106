package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    
    soma := a + b
    sub := a - b
    multi := a * b
    divisao := float64(a) / float64(b)
    resto := a % b
    fmt.Printf("%d\n%d\n%d\n%.2f\n%d\n", soma, sub, multi, divisao, resto)
}