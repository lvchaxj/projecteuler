package main

import "fmt"

func main() {
	var prime [10001]int
	prime[0] = 2
	for i, n := 1, 3; i <= 10000; n += 2 {
		b := true
		for j := 0; j < i; j++ {
			if n%prime[j] == 0 {
				b = false
				break
			}
		}
		if b {
			prime[i] = n
			i++
		}
	}
	fmt.Println(prime)
}
