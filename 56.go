package main

import "fmt"

func main() {
	n := 1
	for a := 2; a < 100; a++ {
		v := []int{1}
		for b := 1; b < 100; b++ {
			x := a % 10
			w := make([]int, len(v))
			c := 0
			for i := 0; i < len(v); i++ {
				c += x * v[i]
				w[i] = c % 10
				c /= 10
			}
			if c > 0 {
				w = append(w, c)
			}

			x = a / 10
			if x > 0 {
				u := make([]int, len(v)+1)
				c = 0
				for i := 0; i < len(v); i++ {
					c += x * v[i]
					u[i+1] = c % 10
					c /= 10
				}
				if c > 0 {
					u = append(u, c)
				}

				c = 0
				for i := 0; i < len(w); i++ {
					c += w[i] + u[i]
					u[i] = c % 10
					c /= 10
				}
				for i := len(w); i < len(u); i++ {
					c += u[i]
					u[i] = c % 10
					c /= 10
				}
				if c > 0 {
					u = append(u, c)
				}
				w = u
			}
			v = w

			fmt.Println(a, b, v)

			m := 0
			for i := 0; i < len(w); i++ {
				m += w[i]
			}

			if m > n {
				n = m
			}
		}
	}
	fmt.Println(n)
}
