package main
import "fmt"
func main() {
    var p, e int
    fmt.Scan(&p, &e)

    forca := 1
    
    for {
        posicao := 0
        salto := forca

        for {
            posicao += salto

            if posicao >= p {
                fmt.Println(forca)
                return
            }

            posicao -= e

            if posicao < 0 {
                break
            }

            if salto > 0 {
                salto -= 10
                if salto < 0 {
                    salto = 0
                }
            }
        }
        forca++
    }
}