package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type SI []rune

var ss, tt []rune
var K [3]SI

func check(c rune) bool {
	if unicode.IsLetter(c) || unicode.IsPunct(c) || unicode.IsSpace(c) || unicode.IsDigit(c) {
		return true
	}
	return false
}

func decrypt(i int) {
	if i == 3 {
		n := 0
		for k := 0; k < len(tt); k++ {
			n += int(tt[k])
		}
		fmt.Println(n)
		return
	}

	for j := 0; j < len(K[i]); j++ {
		for k := i; k < len(ss); k += 3 {
			tt[k] = rune(ss[k]) ^ K[i][j]
		}
		decrypt(i + 1)
	}
}

func main() {
	var s string
	fmt.Scanf("%s", &s)
	t := strings.Split(s, ",")
	ss = make([]rune, len(t))
	tt = make([]rune, len(t))
	for i := 0; i < len(t); i++ {
		v, _ := strconv.ParseInt(t[i], 10, 32)
		ss[i] = rune(v)
	}

	for i := 0; i < 3; i++ {
		for j := rune(0x61); j <= rune(0x7A); j++ {
			b := true
			for k := i; k < len(ss); k += 3 {
				c := rune(ss[k]) ^ j
				if !check(rune(c)) {
					b = false
					break
				}
			}
			if b {
				K[i] = append(K[i], j)
			}
		}
	}

	decrypt(0)
}
