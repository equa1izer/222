package main

import "github.com/01-edu/z01"

func main() {
	arr := [10]int{104, 101, 108, 108, 111, 16, 21, 42}
	PrintMemory(arr)
}

func PrintMemory(arr [10]int) {
	for i, n := range arr {

		s := getHex(n)

		printString(s)

		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		}
	}

	z01.PrintRune('\n')

	for _, n := range arr {
		if n >= 32 {
			z01.PrintRune(rune(n))
		} else {
			z01.PrintRune('.')
		}
	}

	z01.PrintRune('\n')
}

func getHex(n int) string {
	base := "0123456789abcdef"

	res := ""
	for n != 0 {
		i := n % 16
		res = string(base[i]) + res
		n /= 16
	}

	res = extend(res)

	return res
}

func extend(s string) string {
	for len(s) <= 8 {
		if len(s) == 4 {
			s += " "
		}
		s += "0"
	}

	return s + " "
}

func printString(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
