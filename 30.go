package main

import "fmt"

func main() {
	sum := 0
	A := make([]int, 10000000)

	for i := 0; i < 10; i++ {
		A[i] = 1
	}
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			A[i] *= i
		}
	}
	fmt.Println(A[0:10])

	for i := 10; i < 10000000; i++ {
		A[i] = A[i%10] + A[i/10]
		if A[i] == i {
			sum += i
		}
	}

	fmt.Println(sum)

}
