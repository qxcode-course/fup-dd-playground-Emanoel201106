package main

import "fmt"
func main() {
    var num1, num2 int
    var simbolo string
    fmt.Scan(&num1, &num2, &simbolo)
    if simbolo == "+"{
        fmt.Println(num1 + num2) 
    } else if simbolo == "/"{
        fmt.Println(num1 / num2)
    } else if simbolo == "-"{
        fmt.Println(num1 - num2)
    } else {
        fmt.Println(num1 * num2)
    }
    
}
