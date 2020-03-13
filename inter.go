package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		atext, btext := os.Args[1], os.Args[2]
		ftext := ""
		check := false

		for _, a := range atext {
			check = false
			for _, b := range btext {
				if a == b {
					check = true
					for _, f := range ftext {
						if a == f {
							check = false
						}
					}
				}
			}
			if check == true {
				ftext = ftext + string(a)
			}
		}
		for _, l := range ftext {
			z01.PrintRune(rune(l))
		}
	}
	z01.PrintRune('\n')
}
