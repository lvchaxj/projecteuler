package main

import (
	"fmt"
	"math"
)

var P = []int{8, 7, 6, 5, 4, 3, 2, 1}
var max = 0

func permute(i int) {
	if i == len(P)-1 {
		c := 0
		for i := 0; i < len(P); i++ {
			c = c*10 + P[i]
		}

		if c%2 == 0 {
			return
		}
		b := true
		for i := 3; i <= int(math.Sqrt(float64(c))); i += 2 {
			if c%i == 0 {
				b = false
			}
		}
		if b && (c > max) {
			max = c
		}
		return
	}

	for j := i; j < len(P); j++ {
		P[j], P[i] = P[i], P[j]
		permute(i + 1)
		P[j], P[i] = P[i], P[j]
	}
}

func main() {
	permute(0)

	fmt.Println(max)
}
