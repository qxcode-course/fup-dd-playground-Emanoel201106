package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a > b {
		for ; a < b; a++ {
			fmt.Printf("%d", a)
		}
	}
}
