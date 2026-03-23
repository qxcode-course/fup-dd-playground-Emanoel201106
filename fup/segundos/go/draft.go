package main
import "fmt"
func main() {
    var segundos, minuto, hora int
    fmt.Scan(&segundos, &minuto, &hora)
    hora = segundos/3600
    minuto = (segundos % 3600)/60
    segundos = (segundos % 3600) % 60
   
    fmt.Printf("%d:%d:%d\n", hora, minuto, segundos)
}
