package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	count := 0
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	a := Atoi(os.Args[1])

	for i := 0; i <= a; i++ {
		if IsPrime(i) {
			count = count + i
		}
	}
	PrintNbr(count)
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	result := '0'
	if n < 0 {
		z01.PrintRune('-')
		if n/10 != 0 {
			PrintNbr(n / -10)
		}
		for i := 0; i < -(n % 10); i++ {
			result++
		}
		z01.PrintRune(result)
	} else if n == 0 {
		z01.PrintRune('0')
	} else {
		if n/10 != 0 {
			PrintNbr(n / 10)
		}
		for i := 0; i < n%10; i++ {
			result++
		}
		z01.PrintRune(result)
	}
}

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i*i <= nb; i = i + 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func Atoi(s string) int {
	str := 0
	signCount := 0
	addMinus := false
	for i, v := range s {
		if int(rune(v)) > 47 && int(rune(v)) < 58 || v == '+' || v == '-' {
			if (v == '-' || v == '+') && i > 0 {
				return 0
			}
			if v == '-' && i == 0 {
				addMinus = true
			}
			if v == '+' || v == '-' {
				signCount++
				if signCount > 1 {
					return 0
				}
				continue
			}
			a := 0
			for i := '1'; i <= v; i++ {
				a++
			}
			str = str*10 + a
		} else {
			return 0
		}
	}
	if addMinus {
		str = str - (str * 2)
	}
	return str
}
