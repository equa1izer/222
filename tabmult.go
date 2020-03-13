package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) == 2 {

		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			return
		}
		tablmult(n)
	} else {
		z01.PrintRune('\n')
	}

}
func tablmult(nbr int) {
	for i := 1; i <= 9; i++ {
		x := i * nbr
		z01.PrintRune(rune(i + 48))
		printStr(" x ")
		printStr(itoa(nbr))
		printStr(" = ")
		printStr(itoa(x))
		printStr("\n")

	}
}
func itoa(nbr int) string {
	str := ""
	for i := nbr; i != 0; i = i / 10 {
		xx := i % 10
		str = string(rune(xx+48)) + str
	}
	if nbr == 0 {
		str = "0"
	}
	return str

}

func printStr(str string) {
	for _, c := range str {
		z01.PrintRune(c)
	}
}
