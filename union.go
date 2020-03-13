package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) == 3 {
		var array []rune
		for _, c := range os.Args[1] {
			if checkArr(c, array) {
				array = append(array, c)
			}
		}
		for _, c := range os.Args[2] {
			if checkArr(c, array) {
				array = append(array, c)
			}

		}
		for _, c := range array {
			z01.PrintRune(c)
		}

	}
	z01.PrintRune('\n')

}

func checkArr(r rune, arr []rune) bool {
	for _, c := range arr {
		if c == r {
			return false
		}
	}
	return true
}
