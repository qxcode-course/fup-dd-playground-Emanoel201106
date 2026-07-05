package main
import (
    "fmt"
    "strings"
)
//import "strings"
func main() {
    var frase string 
    fmt.Scan(&frase)
    
    partes := strings.Split(frase, "")
    for i := len(partes) - 1; i >= 0; i--{
        fmt.Print(partes[i])
    } 
}
