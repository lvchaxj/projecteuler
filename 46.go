package main

import (
	"fmt"
	"math"
)

const N = 1000000
const K = 6

func main() {
	P := []int{2}
	A := make([]int, N)
	A[1] = 1
	A[2] = 1

	for i := 3; i < N; i += 2 {
		b := true
		for j := 0; j < len(P) && float64(P[j]) <= math.Sqrt(float64(i)); j++ {
			if i%P[j] == 0 {
				b = false
				break
			}
		}
		if b {
			P = append(P, i)
			A[i] = 1
		} else {
			bb := false
			for j := 1; 2*j*j < i; j++ {
				if A[i-2*j*j] == 1 {
					bb = true
					break
				}
			}
			if bb == false {
				fmt.Println(i)
				return
			}
		}
	}
}
