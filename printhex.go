package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		v, err := strconv.Atoi(args[0])
		if err != nil {
			z01.PrintRune('0')
		}
		dec2hex(v)
	} else {
		z01.PrintRune(10)
	}

}

func dec2hex(n int) {
	var arr [100]rune
	i := 0
	for n != 0 {
		temp := n % 16
		if temp < 10 {
			arr[i] = rune(temp + 48)
			i++
		} else {
			arr[i] = rune(temp + 87)
			i++
		}
		n = n / 16
	}
	for j := i - 1; j >= 0; j-- {
		z01.PrintRune(arr[j])
	}
	z01.PrintRune('\n')
}
