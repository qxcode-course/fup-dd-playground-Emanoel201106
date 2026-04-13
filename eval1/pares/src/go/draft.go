package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Println("invalido")
	} else if a != 1{
        for ; a < b; a = a + 2 {
            fmt.Println(a + b) 

        }   
    } 
}
