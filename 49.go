package main

import (
	"fmt"
	"sort"
)

const N = 10000
const K = 6

type S struct {
	v int
	s int
}

type SSlice []S

func (s SSlice) Len() int           { return len(s) }
func (s SSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SSlice) Less(i, j int) bool { return s[i].s < s[j].s || (s[i].s == s[j].s && s[i].v < s[j].v) }

func main() {
	P := []int{2}
	A := []S{}

	for i := 3; i < N; i += 2 {
		b := true
		for j := 0; j < len(P); j++ {
			if i%P[j] == 0 {
				b = false
			}
		}
		if b {
			P = append(P, i)
			if i > 1000 {
				a := []int{i % 10, i / 10 % 10, i / 100 % 10, i / 1000}
				sort.Sort(sort.IntSlice(a))
				A = append(A, S{i, a[0]*1000 + a[1]*100 + a[2]*10 + a[3]})
			}
		}
	}

	sort.Sort(SSlice(A))

	for i := 0; i < len(A); {
		j := 0
		for j = i + 1; j < len(A) && A[i].s == A[j].s; j++ {
		}
		if j-i >= 3 {
		}

		for ii := i; ii < j-1; ii++ {
			for jj := i + 1; jj < j; jj++ {
				u := 2*A[jj].v - A[ii].v
				for kk := i; kk < j; kk++ {
					if kk != ii && kk != jj && A[kk].v == u {
						fmt.Println(A[ii].v, A[jj].v, A[kk].v)
					}
				}
			}
		}

		i = j
	}

}
