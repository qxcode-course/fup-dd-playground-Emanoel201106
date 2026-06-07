package main
import "fmt"
func main() {
    var procurado, n int
    fmt.Scan(&procurado, &n)
    
    contador := 0
    nums := make([]int, n)

    for i := range nums {
        fmt.Scan(&nums[i])
    }
    for _, valor := range nums {
        if valor == procurado {
            contador++
        }
    }
    fmt.Println(contador)
}