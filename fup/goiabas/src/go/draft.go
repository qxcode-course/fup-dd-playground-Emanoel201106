package main
import "fmt"
func main() {
    var cesta, banana, goiaba, manga int
    fmt.Scan(&cesta, &banana, &goiaba, &manga)

    total := banana + goiaba + manga

    if cesta >= total{
        fmt.Println(1)
    } else{
        tempo := total / cesta

         if total % cesta != 0 {
            tempo++
        }
        fmt.Println(tempo)
    }
}
