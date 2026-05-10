package main

import (
	"fmt"
	"strconv"
)
func main() {
    var inicio, fim int
    fmt.Scan(&inicio, &fim)
    for i := inicio; i <= fim; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            zigzag := strconv.Itoa(i)
            zigzag = "zigzag"
            fmt.Println(zigzag)
        } else if i % 5 == 0{
            zag := strconv.Itoa(i)
            zag = "zag"
            fmt.Println(zag)
        }else if i % 3 == 0{
            zig := strconv.Itoa(i)
            zig = "zig"
            fmt.Println(zig)
        } else {
            fmt.Println(i)
        }
        
    }
}
