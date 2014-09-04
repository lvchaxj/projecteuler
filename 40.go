package main

import "fmt"

func main() {
	i, j, k, l, m, n := 2, 10, 10, 1, 1, 1
	for ; k <= 1000000; i++ {
		if i == j {
			j *= 10
			l++
		}
		m += l
		if m >= k {
			v := i
			for ii := m - k; ii > 0; ii-- {
				v /= 10
			}
			n *= v % 10

			k *= 10
		}

	}

	fmt.Println(n)
}
