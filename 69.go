// 69.go
package main

import (
	"fmt"
)

func main() {
	mm, nn := 2, 6
	P := []int{2}
	for n := 3; n < 1000000; n++ {
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

		m := n - k
		if n*mm > m*nn {
			nn, mm = n, m
		}
	}
	fmt.Println(nn)
}
