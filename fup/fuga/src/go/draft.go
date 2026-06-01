package main
import "fmt"
func main() {
    var h, p, f, d int
    fmt.Scan(&h, &p, &f, &d)

    if h >= 0 && h < 16 && p >= 0 && p < 16 && f >= 0 && f < 16 && h != p && h != f && p != 0 {
        for h != f {
            f= (f + d + 16) % 16

            if f ==h {
                fmt.Println("S")
                break
            }
            if f == p {
                fmt.Println("N")
                break
            }
        }
    } else {
        return
    }
}
