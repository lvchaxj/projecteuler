package main

import (
	"fmt"
	"math"
)

const N = 1000000
const K = 11

func prime(v int) bool {
	for j := 3; float64(j) <= math.Sqrt(float64(v)); j += 2 {
		if v%j == 0 {
			return false
		}
	}
	return true
}

func main() {
	P := []int{1, 3, 7, 9}
	Q := make([]int, N)
	Q[0], Q[1], Q[2], Q[3] = 2, 3, 5, 7
	bq, eq := 0, 4

	sum := 0
	for k := K; k > 0 && bq != eq; {
		for _, u := range P {
			v := Q[bq]
			v = v*10 + u
			if prime(v) {
				Q[eq] = v
				eq++
			} else {
				continue
			}

			if u == 9 || u == 1 {
				continue
			}

			b := true
			for j := 100; v > j; j *= 10 {
				if prime(v%j) == false {
					b = false
					break
				}
			}

			if b {
				k--
				sum += v
				fmt.Println(v)
			}
		}

		bq++
	}

	fmt.Println(sum)
}
