package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    nums1 := make([]int, n)
    for i := range nums1 {
        fmt.Scan(&nums1[i])
    }

    var m int
    fmt.Scan(&m)

    nums2 := make([]int, m)
    for i := range nums2 {
        fmt.Scan(&nums2[i])
    }

    usado := make([]bool, len(nums2))
	todos := true
	for _, x := range nums1 {
        encontrou := false
		for j, y := range nums2 {
			if x == y && !usado[j] {
                usado[j] = true
				encontrou = true
                break
			}
		}
        if !encontrou {
            todos = false
            break
        }
	}
	if todos {
		fmt.Println("sim")
	} else {
		fmt.Println("nao")
	}
}
