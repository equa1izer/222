package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		n := 0
		m := 0
		for i := 0; i < len(arg); i++ {
			if arg[i] >= 'A' && arg[i] <= 'Z' {
				n = int(arg[i]) - 64
				for k := 0; k < n; k++ {
					z01.PrintRune(rune(arg[i]))
				}
			} else if arg[i] >= 'a' && arg[i] <= 'z' {
				m = int(arg[i]) - 96
				for j := 0; j < m; j++ {
					z01.PrintRune(rune(arg[i]))
				}
			} else {
				z01.PrintRune(rune(arg[i]))
			}
		}
	}
	z01.PrintRune('\n')
}
