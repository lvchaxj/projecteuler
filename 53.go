// 53.go
package main

import (
	"fmt"
)

func main() {
	n := 0
	for i := 23; i <= 100; i++ {
		v := float64(1)
		for j := 1; j <= i/2; j++ {
			v = v * float64(i+1-j) / float64(j)
			if v > 1000000.0 {
				fmt.Println(i, j, v)
				if i%2 == 0 && j == i/2 {
					n = n + 1
				} else {
					n = n + 2
				}
			}
		}
	}
	fmt.Println(n)
}
