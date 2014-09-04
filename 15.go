package main

import "fmt"

const n int = 20

func main() {
	var DP [n + 1]int
	for j := 0; j <= n; j++ {
		DP[j] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			DP[j] += DP[j-1]
		}
	}

	fmt.Println(DP[n])
}
