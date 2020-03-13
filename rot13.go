package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	str := os.Args[1]
	for _, s := range str {
		if s >= 'a' && s < 'n' || s >= 'A' && s < 'N' {
			z01.PrintRune(s + 13)
		} else if s >= 'n' && s <= 'z' || s >= 'N' && s <= 'Z' {
			z01.PrintRune(s - 13)
		} else {
			z01.PrintRune(s)
		}
	}
	z01.PrintRune('\n')
}
