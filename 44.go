package main

import "fmt"

const N = 10000

func main() {
	A := make([]int, N)
	M := make(map[int]int)

	for i := 1; i < N; i++ {
		A[i] = i * (3*i - 1) / 2
		M[A[i]] = 1
	}

	for i := 1; i < N; i++ {
		for j := i + 1; j < N && 3*j-2 <= A[i]; j++ {
			_, ok1 := M[A[j]-A[i]]
			_, ok2 := M[2*A[j]-A[i]]

			if ok1 && ok2 {
				fmt.Println(A[i])
				return
			}
		}
	}

}
