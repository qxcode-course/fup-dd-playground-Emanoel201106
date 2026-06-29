package main
import "fmt"
func main() {
    var qtd int
    fmt.Scan(&qtd)
    
    colecao := make([]int, qtd)
    for i := range colecao {
        fmt.Scan(&colecao[i])
    }
    
    maior := 0
    cont := 1
    resposta := []int{}

    for i := 1; i < qtd; i++ {
        if colecao[i] == colecao[i-1] {
            cont++
        } else {
            if cont > maior {
                maior = cont
                resposta = []int{colecao[i-1]}
            } else if cont == maior {
                resposta = append(resposta, colecao[i-1])
            }
            cont = 1
        }
    }

    if cont > maior {
        resposta = []int{colecao[qtd-1]}
    } else if cont == maior {
        resposta = append(resposta, colecao[qtd-1])
    }

    fmt.Print("[ ")
    for i := range resposta {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(resposta[i])
    }
    fmt.Println(" ]")
}