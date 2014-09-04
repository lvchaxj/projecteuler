// 71.go
package main

import (
	"fmt"
)

func main() {
	nn, dd := 0, 2
	for d := 3; d <= 1000000; d++ {
		if d == 7 {
			continue
		}
		n := 3 * d / 7
		a, b := n, d
		for b != 0 {
			a, b = b, a%b
		}
		if a != 1 {
			continue
		}
		if n*dd > d*nn {
			nn, dd = n, d
		}
	}
	fmt.Println(nn)
}
