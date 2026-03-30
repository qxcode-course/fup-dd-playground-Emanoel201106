package main
import "fmt"
func main() {
    var num1, num2 int
    fmt.Scan(&num1, &num2)
    diferença := num1 - num2
    if diferença < 0 {
        diferença = diferença * (-1)
    }
    fmt.Println(diferença)
}
