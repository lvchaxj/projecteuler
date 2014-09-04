package main

import "fmt"

const N = 10000000

func main() {
	A := make([]int, N)

	A[0] = 1
	for i := 1; i < 10; i++ {
		A[i] = i * A[i-1]
	}

	sum := 0
	for i := 10; i < N; i++ {
		A[i] = A[i/10] + A[i%10]
		if i == A[i] {
			sum += i
		}
	}

	fmt.Println(sum)
}
