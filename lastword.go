package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		mas := []rune(arg)
		var s string
		t := false

		for i := len(arg) - 1; i >= 0; i-- {
			if mas[i] == ' ' && t {
				break
			} else if mas[i] != ' ' {
				s = string(mas[i]) + s
				t = true
			}
		}
		for _, j := range s {
			z01.PrintRune(j)
		}
	}
	z01.PrintRune('\n')
}
