package main
import "fmt"
func main() {
    var num1, num2 int
    fmt.Scan(&num1, &num2)

    if num1 % 3 == 0 && num2 % 3 == 0 || num1 % 5 == 0 && num2 % 5 == 0{
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}