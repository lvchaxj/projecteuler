package main

import "fmt"

func main() {
	primes := []int{2, 3}

	for i, n := 3, 3; ; i++ {
		n += i
		nn := n
		m := 1
		for j := 0; nn > 1 && j < len(primes); j++ {
			k := 1
			for ; nn > 1 && nn%primes[j] == 0; nn /= primes[j] {
				k++
			}
			m *= k
		}

		if nn > 1 {
			primes = append(primes, nn)
			m *= 2
		}

		if m > 500 {
			fmt.Println(n)
			break
		}
	}
}
