package main

import "fmt"
import "math"

func prime(n int) bool {
	if n%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	a, m, n := 1, 0, 1
	for i := 3; ; i += 2 {
		for j := 0; j < 4; j++ {
			a += i - 1
			if prime(a) {
				m++
			}
			n++
		}
		if float64(m)/float64(n) < 0.1 {
			fmt.Println(i)
			break
		}
	}
}
