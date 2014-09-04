package main

import "fmt"

func main() {
	n := 1000000
	m := 0
	V := make([]int, 10)

	p := 1
	for i := 1; i < 10; i++ {
		p *= i
		V[i] = i
	}

	for i := 9; i > 0; i-- {
		j := (n - 1) / p
		m = m*10 + V[j]

		V = append(V[0:j], V[j+1:]...)

		n -= j * p
		p /= i
	}

	m = m*10 + V[0]

	fmt.Println(m)
}
