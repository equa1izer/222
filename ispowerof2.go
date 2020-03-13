package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		value, err := strconv.Atoi(os.Args[1])
		if err != nil {
			for _, l := range err.Error() {
				z01.PrintRune(l)
			}
			return
		}
		if value%2 == 0 && value >= 0 {
			for _, l := range "true" {
				z01.PrintRune(l)
			}
		} else {
			for _, l := range "false" {
				z01.PrintRune(l)
			}
		}

	}
}
