package main
import "fmt"
func main() {
    var esq1, esq2, esq3 int
    fmt.Scan(&esq1, &esq2, &esq3)
    if (esq1 > esq2 && esq1 < esq3) || (esq1 < esq2 && esq1 > esq3){
        fmt.Println(esq1)
    } else if (esq2 > esq1 && esq2 < esq3) || (esq2 < esq1 && esq2 > esq3){
        fmt.Println(esq2)
    } else {
        fmt.Println(esq3)
    }
}
