package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)

	words := strings.Split(s, ",")
	sort.Sort(sort.StringSlice(words))

	fmt.Println(words[0])
	var sum int64
	for i, s := range words {
		var n int64
		for j := 1; j < len(s)-1; j++ {
			n += int64(s[j] - 'A' + 1)
		}
		sum += n * int64(i+1)
	}

	fmt.Println(sum)
}
