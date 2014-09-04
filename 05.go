package main

import "fmt"
import "math"

func main() {
	m := make(map[int]int)
	m[2] = 1

	n := 0
	for i := 3; i <= 10; i++ {
		n = i
		for a, b := range m {
			c := 0
			for ; n > 1 && n%a == 0; n /= a {
				c++
			}
			if c > b {
				m[a] = c
			}
		}
		if n > 1 {
			m[n] = 1
		}
	}

	n = 1
	for a, b := range m {
		n *= int(math.Pow(float64(a), float64(b)))
	}

	fmt.Println(n)
}
