package main

import "fmt"

func main() {
	n := 0
	a, b := []int{1}, []int{1}
	for i := 0; i < 1000; i++ {
		// a = a+b
		d := 0
		for j := 0; j < len(b); j++ {
			d += a[j] + b[j]
			a[j] = d % 10
			d /= 10
		}
		for j := len(b); j < len(a); j++ {
			d += a[j]
			a[j] = d % 10
			d /= 10
		}
		if d > 0 {
			a = append(a, d)
		}

		// c = a
		c := make([]int, len(a))
		copy(c, a)

		// a = a+b
		d = 0
		for j := 0; j < len(b); j++ {
			d += a[j] + b[j]
			a[j] = d % 10
			d /= 10
		}
		for j := len(b); j < len(a); j++ {
			d += a[j]
			a[j] = d % 10
			d /= 10
		}
		if d > 0 {
			a = append(a, d)
		}

		// b = c
		b = c

		if len(a) > len(b) {
			n++
		}
	}

	fmt.Println(n)
}
