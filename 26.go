package main

import "fmt"

func main() {
	var N [1000]int
	for i := 2; i < 1000; i++ {
		M := make([]int, i)
		n := 0
		k := 1
		for k > 0 {
			for k < i {
				n++
				k *= 10
			}
			k %= i
			if M[k] > 0 {
				break
			} else {
				M[k] = n
			}
		}
		if k > 0 {
			N[i] = n - M[k]
		}
	}

	k := 0
	for i, n := range N {
		if n > N[k] {
			k = i
		}
	}

	fmt.Println(k)
}
