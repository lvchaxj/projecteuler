// 70.go
package main

import (
	"fmt"
)

func main() {
	mm, nn := 79180, 87109
	P := []int{2}
	for n := 3; n < 10000000; n++ {
		k := 0
		b := true
		for i := 0; i < len(P) && P[i]*P[i] <= n; i++ {
			if n%P[i] == 0 {
				k = k + (n-k)/P[i]
				if P[i]*P[i] != n {
					k = k + (n-k)*P[i]/n
				}
				b = false
			}
		}
		if b {
			P = append(P, n)
			k = 1
		}
		if k > 0 {
			m := n - k
			var X, Y [10]int
			for i := m; i > 0; i /= 10 {
				X[i%10]++
			}
			for i := n; i > 0; i /= 10 {
				Y[i%10]++
			}
			b = true
			for i := 0; i < 10; i++ {
				if X[i] != Y[i] {
					b = false
					break
				}
			}
			if b {
				if n*mm < m*nn {
					nn, mm = n, m
				}
			}
		}
	}
	fmt.Println(nn)
}
