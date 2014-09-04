package main

import "fmt"

const max int = 1000000

func main() {
	N := 1

	var m, queue [max + 1]int
	m[1] = 1
	queue[0] = 1
	qb, qe := 0, 1

	for qb < qe {
		v := queue[qb]
		qb++

		a := 2 * v
		if a <= max {
			if m[a] == 0 {
				m[a] = m[v] + 1
				queue[qe] = a
				qe++
			}
		}

		if (v-1)%3 == 0 && (v-1)%6 != 0 {
			b := (v - 1) / 3
			if m[b] == 0 {
				m[b] = m[v] + 1
				queue[qe] = b
				qe++
			}
		}
	}

	for i := 1; i <= max; i++ {
		if m[i] > 0 {
			if m[i] > m[N] {
				N = i
			}
		} else {
			for j, n := i, 1; ; n++ {
				if j%2 == 0 {
					j /= 2
				} else {
					j = 3*j + 1
				}
				if j < max && m[j] > 0 {
					m[i] = m[j] + n
					break
				}
			}
			if m[i] > m[N] {
				N = i
			}
		}
	}
	fmt.Println(N)
}
