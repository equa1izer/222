package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 0 {
		arg := os.Args[1]
		mas := []rune(arg)
		search(mas)
	}
	z01.PrintRune('\n')
}

func firstWord(mas []rune) int {
	var s string
	c := 0
	for _, i := range mas {
		if i != ' ' {
			s = s + string(i)
		} else if s != "" {
			break
		}
		c++
	}
	return c
}

func search(mas []rune) {
	var s string
	c := 0
	c = firstWord(mas)
	for i := c + 1; i < len(mas); i++ {
		if mas[i] != ' ' {
			s = s + string(mas[i])
		}
		if mas[i-1] != ' ' && (mas[i] == ' ' || i == len(mas)-1) {
			s = s + " "
		}
	}
	for _, i := range s {
		z01.PrintRune(i)
	}
	for i := 0; i < c; i++ {
		if mas[i] != ' ' {
			z01.PrintRune(mas[i])
		}
	}
}
