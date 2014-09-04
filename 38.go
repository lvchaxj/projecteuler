package main

import "fmt"

var P = [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
var max = 918273645

func permute(i int) {
	if i == len(P) {
		a := 1000*P[0] + 100*P[1] + 10*P[2] + P[3]
		b := 10000*P[4] + 1000*P[5] + 100*P[6] + 10*P[7] + P[8]
		c := 100000000*P[0] + 10000000*P[1] + 1000000*P[2] + 100000*P[3] + 10000*P[4] + 1000*P[5] + 100*P[6] + 10*P[7] + P[8]

		if a*2 == b && c > max {
			max = c
		}
		return
	}
	for j := i; j < len(P); j++ {
		P[j], P[i] = P[i], P[j]
		permute(i + 1)
		P[j], P[i] = P[i], P[j]
	}
}

func main() {
	permute(1)

	fmt.Println(max)
}
