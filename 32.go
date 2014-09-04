package main

import "fmt"

var A = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var M = make([]int, 10000)

func perm(i int, sum *int) {
	if i == len(A)-1 {
		a, b := 0, 0
		c := 1000*A[5] + 100*A[6] + 10*A[7] + A[8]
		if M[c] == 1 {
			return
		}

		a = A[0]
		b = 1000*A[1] + 100*A[2] + 10*A[3] + A[4]
		if a*b == c {
			*sum += c
			M[c] = 1
			return
		}

		a = 10*A[0] + A[1]
		b = 100*A[2] + 10*A[3] + A[4]
		if a*b == c {
			*sum += c
			M[c] = 1
			return
		}

		a = 100*A[0] + 10*A[1] + A[2]
		b = 10*A[3] + A[4]
		if a*b == c {
			*sum += c
			M[c] = 1
			return
		}

		a = 1000*A[0] + 100*A[1] + 10*A[2] + A[3]
		b = A[4]
		if a*b == c {
			*sum += c
			M[c] = 1
			return
		}

		return
	}
	for j := i; j < len(A); j++ {
		A[i], A[j] = A[j], A[i]
		perm(i+1, sum)
		A[i], A[j] = A[j], A[i]
	}
}

func main() {
	sum := 0
	perm(0, &sum)

	fmt.Println(sum)
}
