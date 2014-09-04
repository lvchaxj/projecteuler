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
		}
	}
	P = append(P, N)

	n := 4
	m := 10
	i := 4
	for k := 1; k < K; k++ {
		for m *= 10; P[i] < m; i++ {
			v := P[i]
			b := true
			for j := 0; j < k; j++ {
				v = v/(m/10) + v%(m/10)*10
				if A[v] == 0 {
					b = false
					break
				}
			}
			if b {
				n++
			}
		}
	}

	fmt.Println(n)
}
