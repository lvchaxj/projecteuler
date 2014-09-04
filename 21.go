package main

import "fmt"

const N = 10000

func main() {
	P := []int{2}
	var A [N + 1]int
	A[2] = 1

	for i := 3; i < N; i++ {
		n := i
		a := 1
		for _, p := range P {
			b := 1
			for c := 1; n%p == 0; n /= p {
				c *= p
				b += c
			}
			a *= b
		}
		if n > 1 {
			P = append(P, n)
			a *= (1 + n)
		}
		A[i] = a - i
	}

	sum := 0
	for i := 2; i < N; i++ {
		if A[i] < i && A[A[i]] == i {
			sum += (i + A[i])
		}
	}

	fmt.Println(sum)
}
