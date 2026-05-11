package main

import "fmt"

func main() {
	var N, X, Y int
	var C string
	var S int
	fmt.Scan(&N, &X, &Y, &C, &S)
	if C == "U" {
		Y = (Y - S%N + N) % N
	} else if C == "D" {
		Y = (Y + S) % N
	} else if C == "L" {
		X = (X - S%N + N) % N
	} else {
		X = (X + S) % N
	}
	fmt.Println(X, Y)
}
