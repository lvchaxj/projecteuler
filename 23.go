package main

import "fmt"

const N = 28123

func main() {
	P := []int{2}
	var A [N + 1]int
	var B [N + 1]int
	e := 0

	for i := 3; i <= N; i++ {
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
		if (a - i) > i {
			A[i] = 1
			B[e] = i
			e++
		}
	}

	sum := (1 + 23) * 23 / 2
	for i := 24; i <= N; i++ {
		f := false
		for j := 0; j < e && B[j] < i; j++ {
			if A[i-B[j]] == 1 {
				f = true
				break
			}
		}
		if f == false {
			sum += i
		}
	}

	fmt.Println(sum)
}
