package main

import "fmt"

type Pessoa struct {
    nome string
    idade int
    sexo string
}

func main() {
	var n int
    fmt.Scan(&n)

    pessoas := make([]Pessoa, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&pessoas[i].nome, &pessoas[i].idade, &pessoas[i].sexo)
    }
    
    indiceMaior := -1

    for i := 0; i < n; i++ {
        if pessoas[i].sexo == "f" {
            if indiceMaior == -1 || pessoas[i].idade > pessoas[indiceMaior].idade {
                indiceMaior = i
            }
        }
    }

    if indiceMaior == -1 {
        fmt.Println("nao tem mulher")
    } else {
        fmt.Println(pessoas[indiceMaior].nome)
    }
}
