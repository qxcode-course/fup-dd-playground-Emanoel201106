package main

import "fmt"

func main() {
	var valor, chute1, chute2 int
	fmt.Scan(&valor, &chute1, &chute2)

	diferença1 := valor - chute1
    if diferença1 < 0 {
        diferença1 = -diferença1
    }
	diferença2 := valor - chute2
    if diferença2 < 0 {
        diferença2 = -diferença2
    }

	if diferença1 < diferença2 {
		fmt.Println("primeiro")
	} else if diferença2 < diferença1 {
		fmt.Println("segundo")
	} else {
		fmt.Println("empate")
	}
}
