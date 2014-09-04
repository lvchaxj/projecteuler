package main

import (
	"fmt"
	"math"
)

func main() {
	var P = make([]int, 1001)

	for i := 1; i < 1000; i++ {
		for j := i; i+j < 1000; j++ {
			k := int(math.Sqrt(float64(i*i + j*j)))
			if i+j+k <= 1000 && i*i+j*j == k*k {
				P[i+j+k]++
			}
		}
	}

	max := 0
	for i := 1; i <= 1000; i++ {
		if P[i] > P[max] {
			max = i
		}
	}

	fmt.Println(max)
}
