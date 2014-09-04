// 69.go
package main

import (
	"fmt"
)

func main() {
	m := 1
	P := []int{2}
	for n := 3; n <= 1000000; n++ {
		k := 0
		b := true
		for i := 0; i < len(P); i++ {
			if n%P[i] == 0 {
				k = k + (n-k)/P[i]
				b = false
			}
		}
		if b {
			P = append(P, n)
			k = 1
		}

		m += n - k
	}
	fmt.Println(m)
}
