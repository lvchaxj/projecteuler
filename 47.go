package main

import (
	"fmt"
	//"math"
)

const N = 1000000
const K = 6

func main() {
	P := []int{2}

	k := 0
	for i := 3; i < N; i++ {
		b := true
		m := 0
		fmt.Printf("%d: ", i)
		for j := 0; j < len(P); j++ {
			if i%P[j] == 0 {
				fmt.Printf("%d ", P[j])
				m++
				b = false
			}
		}
		fmt.Println()
		if b {
			P = append(P, i)
		}

		if m == 4 {
			k++
			if k == 4 {
				fmt.Println(i)
				return
			}
		} else {
			k = 0
		}
	}
}
