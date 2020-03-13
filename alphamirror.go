package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	var str string
	t := false
	if len(arg) == 0 {
		z01.PrintRune('\n')
		t = true
		return
	}
	if !t {
		str = arg[0]
	}
	alphamirror(str)
}

func alphamirror(s string) {
	var str string
	x := 0
	for _, i := range s {
		if i >= 'a' && i <= 'm' {
			for j := 'a'; j <= i; j++ {
				x++
			}
			i = 123 - rune(x+48)
			str = str + string(i+48)
			x = 0
		} else if i >= 'n' && i <= 'z' {
			for j := 'n'; j <= i; j++ {
				x++
			}
			i = 110 - rune(x+48)
			str = str + string(i+48)
			x = 0
		} else if i >= 'A' && i <= 'M' {
			for j := 'A'; j <= i; j++ {
				x++
			}
			i = 91 - rune(x+48)
			str = str + string(i+48)
			x = 0
		} else if i >= 'N' && i <= 'Z' {
			for j := 'N'; j <= i; j++ {
				x++
			}
			i = 78 - rune(x+48)
			str = str + string(i+48)
			x = 0
		} else {
			str = str + string(i)
		}
	}
	for _, i := range str {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
