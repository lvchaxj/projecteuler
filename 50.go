package main

import (
	"fmt"
	"math"
)

const N = 100

func main() {
	P := []int{2}
	Q := []int{2}
	A := make([]int, N+1)
	A[2] = 1

	for i := 3; i < N; i += 2 {
		b := true
		for j := 0; j < len(P) && P[j] <= int(math.Sqrt(float64(i))); j++ {
			if i%P[j] == 0 {
				b = false
			}
		}
		if b {
			P = append(P, i)
			Q = append(Q, Q[len(Q)-1]+i)
			A[i] = 1
		}
	}

	fmt.Println(P)
	fmt.Println(Q)

}
