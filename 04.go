package main

import "fmt"

func palindrome(n int) bool {
	m := 0
	for nn := n; nn > 0; nn /= 10 {
		m = m*10 + nn%10
	}
	return m == n
}

func main() {
	max, n := 0, 0

	for i := 999; i >= 100; i-- {
		for j := 999; j >= i; j-- {
			n = i * j
			if palindrome(n) && n > max {
				max = n
			}
		}
	}

	fmt.Println(max)
}
