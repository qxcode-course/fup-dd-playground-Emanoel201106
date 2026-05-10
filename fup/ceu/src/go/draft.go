package main

import (
	"fmt"
	"strconv"
)
func main() {
    var pedra int
    fmt.Print("[ ")
    fmt.Scan(&pedra)
    n := 10
    for i := 0; i <= n; i++ {
        if i == pedra{
            continue
        } 
        if i == n{
            ceu := strconv.Itoa(n)
            ceu = "ceu "
            fmt.Print(ceu)
            continue
        }
        fmt.Printf("%d ", i)
    }
    fmt.Println("]")
}
