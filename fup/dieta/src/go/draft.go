package main
import "fmt"
func main() {
    var dias int
    fmt.Scan(&dias)
    soma := 0
    var kcal []int = make([]int, dias)
    for i := range kcal {
        fmt.Scan(&kcal[i])
        soma += kcal[i]
    }
    
    media := float64(soma) / float64(dias)
    fmt.Printf("%.1f\n", media)
}
