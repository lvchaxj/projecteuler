package main

import "fmt"

func main() {
	var m [20][20]int
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			fmt.Scanf("%d", &m[i][j])
		}
	}

	max, v := 0, 0
	for i := 0; i < 20; i++ {
		for j := 0; j < 17; j++ {
			v = m[i][j] * m[i][j+1] * m[i][j+2] * m[i][j+3]
			if v > max {
				max = v
			}
			v = m[j][i] * m[j+1][i] * m[j+2][i] * m[j+3][i]
			if v > max {
				max = v
			}
		}
	}

	for i := 0; i < 17; i++ {
		for j := 0; j < 17; j++ {
			v = m[i][j] * m[i+1][j+1] * m[i+2][j+2] * m[i+3][j+3]
			if v > max {
				max = v
			}
			v = m[i][j+3] * m[i+1][j+2] * m[i+2][j+1] * m[i+3][j]
			if v > max {
				max = v
			}
		}
	}

	fmt.Println(max)
}
