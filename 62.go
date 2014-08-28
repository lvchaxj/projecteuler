package main

import "fmt"

func main() {
	var M = make(map[int64]int)
	var N = make(map[int64]int64)

	for i := int64(1); ; i++ {
		v := i * i * i

		var D [10]int
		for v > 0 {
			D[v%10]++
			v /= int64(10)
		}
		v = 0
		for i := 1; i < 10; i++ {
			for j := 0; j < D[i]; j++ {
				v = v*10 + int64(i)
			}
		}
		for j := 0; j < D[0]; j++ {
			v *= 10
		}

		if _, ok := M[v]; ok {
			M[v]++
			if M[v] == 5 {
				fmt.Println(N[v] * N[v] * N[v])
				return
			}
		} else {
			M[v] = 1
			N[v] = i
		}
	}

}
