package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	defer z01.PrintRune('\n')
	neg := "-1"
	if len(os.Args) == 3 {
		a := os.Args[1]
		b := os.Args[2]
		if a > b {
			z01.PrintRune(49)
		} else if a == b {
			z01.PrintRune(48)
		} else {
			for _, r := range neg {
				z01.PrintRune(r)
			}
		}
	}
}
