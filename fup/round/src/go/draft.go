package main
import "fmt"
import "math"

func ceil(num float64) int{
    resto := math.Mod(num, 1)
    return int((num - resto) + 1) 
}

func floor(num float64) int{
    resto := math.Mod(num, 1)
    return int((num - resto))
}

func round(num float64) int{
    resto := math.Mod(num, 1)
    if resto >= 0.5 {
        return ceil(num)
    } else {
        return floor(num)
    }
}

func main() {
    var operação string
    var num float64
   
    fmt.Scan(&operação, &num)
    var resultado int
    if operação == "c" {
        resultado = ceil(num)
    } else if operação == "f" {
        resultado = floor(num)
    } else if operação == "r" {
        resultado = round(num)
    }
    fmt.Println(resultado)
}
