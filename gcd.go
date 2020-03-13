package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 2 {
		a, _ := strconv.Atoi(args[0])
		b, _ := strconv.Atoi(args[1])
		fmt.Println(gcd(a, b))
	} else {
		z01.PrintRune(10)
	}

}
func gcd(a, b int) int {
	var bgcd func(a, b, res int) int

	bgcd = func(a, b, res int) int {
		switch {
		case a == b:
			return res * a
		case a%2 == 0 && b%2 == 0:
			return bgcd(a/2, b/2, 2*res)
		case a%2 == 0:
			return bgcd(a/2, b, res)
		case b%2 == 0:
			return bgcd(a, b/2, res)
		case a > b:
			return bgcd(a-b, b, res)
		default:
			return bgcd(a, b-a, res)
		}
	}

	return bgcd(a, b, 1)
}
