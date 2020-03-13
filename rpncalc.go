package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Error")
		return
	}

	str := trimWhitespaces(leaveOneSpace(args[0]))
	rpn := strings.Split(str, " ")

	if !isValid(rpn) {
		fmt.Println("Error")
		return
	}

	for i := 0; i < len(rpn); i++ {
		arg := rpn[i]

		if arg == "+" || arg == "-" || arg == "/" || arg == "*" || arg == "%" {

			if i < 2 {
				fmt.Println("Error")
				return
			}

			newNode, err := resolve(rpn[i-2], rpn[i-1], arg)

			if err {
				fmt.Println("Error")
				return
			}

			rpn[i-2] = newNode

			rpn = append(rpn[:i-1], rpn[i+1:]...)
			i -= 3
		}
	}

	fmt.Println(rpn[0])
}

func isValid(rpn []string) bool {
	balance := 0
	for _, s := range rpn {
		if s == "+" || s == "-" || s == "/" || s == "*" || s == "%" {
			balance--
		} else {
			balance++
		}
	}

	if balance == 1 {
		return true
	}

	return false
}

func resolve(s1, s2, sign string) (string, bool) {

	n1, err := strconv.Atoi(s1)

	if err != nil {
		return "", true
	}

	n2, err := strconv.Atoi(s2)
	if err != nil {
		return "", true
	}

	res := 0

	switch sign {
	case "+":
		res = n1 + n2
	case "-":
		res = n1 - n2
	case "*":
		res = n1 * n2
	case "/":
		res = n1 / n2
	case "%":
		res = n1 % n2
	default:
		return "", true
	}

	return itoa(res), false
}

func itoa(x int) string {
	r := ""

	if x < 0 {

		for x != 0 {
			r = string('0'-x%10) + r
			x /= 10
		}

		return "-" + r
	}

	for x != 0 {
		r = string('0'+x%10) + r
		x /= 10
	}

	return r
}

func trimWhitespaces(s string) string {
	if len(s) == 0 {
		return s
	}

	if s[0] == ' ' {
		return trimWhitespaces(s[1:])
	}

	if s[len(s)-1] == ' ' {
		return trimWhitespaces(s[:len(s)-1])
	}

	return s
}

func leaveOneSpace(s string) string {
	for i := 1; i < len(s); i++ {
		if s[i-1] == ' ' && s[i] == ' ' {
			s = s[:i-1] + s[i:]
			i--
		}
	}

	return s
}
