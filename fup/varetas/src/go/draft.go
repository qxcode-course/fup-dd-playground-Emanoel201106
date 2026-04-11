package main
import "fmt"
func main() {
    var vare1, vare2, vare3 int
    fmt.Scan(&vare1, &vare2, &vare3)
    if vare1 < vare2 + vare3 && vare2 < vare1 + vare3 && vare3 < vare1 + vare2 {
       fmt.Println("True") 
    } else {
        fmt.Println("False")
    }
}
