package main

import "fmt"

func main() {
	var M = []int{0, 1, 2, 5, 10, 20, 50, 100, 200}
	var DP [9][201]int

	for j := 0; j <= 200; j++ {
		DP[1][j] = 1
	}
	for i := 2; i < len(M); i++ {
		DP[i][0] = 1
		for j := 1; j <= 200; j++ {
			DP[i][j] = DP[i-1][j]
			for jj := j - M[i]; jj >= 0; jj -= M[i] {
				DP[i][j] += DP[i-1][jj]
			}
		}
	}

	fmt.Println(DP[8][200])
}
