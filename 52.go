// 52.go
package main

import (
	"fmt"
)

func main() {
	for n := 1; ; n++ {
		var D [10]int
		for nn := n; nn > 0; nn /= 10 {
			D[nn%10]++
		}

		b := true
		for i := 2; b && i < 6; i++ {
			var E [10]int
			for nn := i * n; nn > 0; nn /= 10 {
				E[nn%10]++
			}
			for j := 0; j < 10; j++ {
				if E[j] != D[j] {
					b = false
					break
				}
			}
		}

		if b {
			fmt.Println(n)
			return
		}
	}
}
