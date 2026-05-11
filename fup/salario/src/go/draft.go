package main
import "fmt"
func main() {
    var salario float64
    fmt.Scan(&salario)
    if salario <= 1000 {
        novosalario := (0.2 * salario) + salario
        fmt.Printf("%.2f\n",novosalario)
    } else if salario <= 1500 {
        novosalario := (0.15 * salario) + salario
        fmt.Printf("%.2f\n", novosalario)
    } else if salario <= 2000 {
        novosalario := (0.1 * salario) + salario
        fmt.Printf("%.2f\n", novosalario)
    } else if salario > 2000{
        novosalario := (0.05 * salario) + salario
        fmt.Printf("%.2f\n", novosalario)
    }
    
}
