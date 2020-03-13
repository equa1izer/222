package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args
	if len(arg) != 2 {
		fmt.Println()
		return
	}
	numb, _ := strconv.Atoi(arg[1])
	fprime(numb)
}

func fprime(n int) {
	if n <= 1 {
		fmt.Println()
		return
	}
	a := 2
	for a <= n {
		if n%a == 0 {
			n = n / a
			if a <= n {
				fmt.Print(a, " * ")
			} else {
				fmt.Print(a, "\n")
			}
		} else {
			a++
		}
	}
}
