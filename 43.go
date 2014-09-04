package main

import (
	"fmt"
)

var P = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var Q = []int{0, 2, 3, 5, 7, 11, 13, 17}

var sum = 0

func permute(i int) {
	if i == len(P)-1 {
		c := 0
		b := true
		for i := 1; i < 8; i++ {
			c = 100*P[i] + 10*P[i+1] + P[i+2]
			if c%Q[i] != 0 {
				b = false
				break
			}
		}
		if b {
			c = 0
			for i := 0; i < len(P); i++ {
				c = c*10 + P[i]
			}
			sum += c
			fmt.Println(c)
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

	fmt.Println(sum)
}
