// 55.go
package main

import (
	"fmt"
)

func main() {
	n := 0
	for i := 1; i <= 10000; i++ {
		v := []int{}
		ii := i
		for ii > 0 {
			v = append(v, ii%10)
			ii /= 10
		}

		for k := 0; k < 50; k++ {
			w := make([]int, len(v))
			c := 0
			for j := 0; j < len(v); j++ {
				c += v[j] + v[len(v)-1-j]
				w[j] = c % 10
				c /= 10
			}
			if c > 0 {
				w = append(w, c)
			}

			v = w
			b := true
			for j := 0; j <= len(v)-1-j; j++ {
				if v[j] != v[len(v)-1-j] {
					b = false
					break
				}
			}

			if b {
				n++
				break
			}
		}
	}
	fmt.Println(10000 - n)
}
