package main

import (
	"fmt"
	"math"
	"os"
)

const n int = 6

var N, M, V, W [n]int

func cyclic(i int) {
	if i == n {
		if V[i-1]%100 == V[0]/100 {
			fmt.Println(V)
			sum := 0
			for j := 0; j < n; j++ {
				sum += V[j]
			}
			fmt.Println(sum)
			os.Exit(0)
		}
		return
	}

	for k := 0; k < n; k++ {
		if W[k] == 1 {
			continue
		}
		W[k] = 1
		a, b := k+1, 1-k
		for M[i] = N[i]; ; M[i]++ {
			V[i] = M[i] * (M[i]*a + b) / 2
			if V[i] > 9999 {
				break
			}
			if i == 0 || V[i]/100 == V[i-1]%100 {
				cyclic(i + 1)
			}
		}
		W[k] = 0
	}
}

func main() {
	for i := 0; i < n; i++ {
		a, b := float64(i+1), float64(1-i)
		d := math.Sqrt(b*b + 4*a*2000)
		N[i] = int(math.Ceil((d - b) / 2 / a))
	}
	cyclic(0)
}
