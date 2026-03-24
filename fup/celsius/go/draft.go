package main
import "fmt"
func main() {
    var celsius, fahrenheit float64
    fmt.Scan(&celsius, &fahrenheit)
    fahrenheit = 1.8 * celsius + 32
    fmt.Printf("%.6f\n", fahrenheit)
}
