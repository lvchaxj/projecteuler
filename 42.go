package main

import (
	"fmt"
	"strings"
)

func main() {
	A := make([]int, 5051)
	for i := 1; i <= 100; i++ {
		A[i*(i+1)/2] = 1
	}

	s := ""
	fmt.Scanf("%s", &s)
	words := strings.Split(s, ",")

	m := 0
	for _, w := range words {
		n := 0
		for i := 1; i < len(w)-1; i++ {
			n += int(w[i] - 'A' + 1)
		}

		if A[n] == 1 {
			m++
		}
	}

	fmt.Println(m)

}
