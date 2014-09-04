package main

import "fmt"

const N = 15

func main() {
	var DP [N][N]int
	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			fmt.Scanf("%d", &DP[i][j])
		}
	}

	for i := 13; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if DP[i+1][j] > DP[i+1][j+1] {
				DP[i][j] += DP[i+1][j]
			} else {
				DP[i][j] += DP[i+1][j+1]
			}
		}
	}

	fmt.Println(DP[0][0])
}
