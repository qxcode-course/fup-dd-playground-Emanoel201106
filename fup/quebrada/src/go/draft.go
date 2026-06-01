package main
import "fmt"
func main() {
    var n1, n2 int
    fmt.Scan(&n1, &n2)

    resultadoi := n1 / n2
    resto := n1 % n2
    resultadof := float64(n1) / float64(n2)
    fmt.Printf("%d\n%d\n%.2f\n", resultadoi, resto, resultadof)
}