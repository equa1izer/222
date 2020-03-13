package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		str := os.Args[1]

		num := map[rune]int{}

		for _, i := range str {
			num[i]++
		}
		res := ""
		for key, count := range num {
			if isUniq(num, key, count) {
				res = "true"
			} else {
				toPrint("false")
				return
			}
		}
		toPrint(res)
	}
	z01.PrintRune('\n')
}
func toPrint(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
}
func isUniq(m map[rune]int, key rune, count int) bool {
	for s, num := range m {
		if s != key {
			if num == count {
				return false
			}
		}
	}
	return true
}
