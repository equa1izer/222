package main

//last/firstruneprog
func main() {
	fmt.Println(FirstRune("firstruna"))
}
func FirstRune(str string) (res rune) {
	count := 0
	for range str {
		count++
	}
	return rune(str[count-1])
}

//rot13
func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		for _, r := range os.Args[1] {
			if (r >= 65 && r <= 77) || (r >= 97 && r <= 109) {
				z01.PrintRune(r + 13)
			} else if (r >= 78 && r <= 92) || (r >= 110 && r <= 122) {
				z01.PrintRune(r - 13)
			} else {
				z01.PrintRune(r)
			}
		}
	}
}

//lastword
func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		var start, end int
		var slice string
		str := os.Args[1]
		for i := len(str) - 1; i != 0; i-- {
			if end == 0 && str[i] != ' ' {
				end = i + 1
			}
			if end > 0 && str[i] == ' ' {
				start = i + 1
			}
		}
		slice = str[start:end]
		for _, r := range slice {
			z01.PrintRune(r)
		}
	}
}

//robotorigin
func main() {
	defer z01.PrintRune('\n')
	result := "false"
	var way int
	if len(os.Args) == 2 {
		for _, r := range os.Args[1] {
			if r == 'U' || r == 'R' {
				way++
			}
			if r == 'D' || r == 'L' {
				way--
			}
		}
		if way == 0 {
			result = "true"
		}
	}
	for _, r := range result {
		z01.PrintRune(r)
	}
}

//tabmult
func main() {
	var res string
	if len(os.Args) == 2 {
		num, _ := strconv.Atoi(os.Args[1])
		for i := 1; i <= 9; i++ {
			res = res + iToA(i) + " x " + iToA(num) + " = " + iToA(num*i) + "\n"
		}
		for _, r := range res {
			z01.PrintRune(r)
		}
	} else{
		z01.PrintRune('\n')
	}

}
func iToA(num int) (str string) {
	for num > 0 {
		str = string(num%10+48) + str
		num = num / 10
	}
	return str
}

alphamirror
func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		for _, r := range os.Args[1] {
			if r >= 65 && r <= 90 {
				z01.PrintRune(90 - (r - 65))
			} else if r >= 97 && r <= 122 {
				z01.PrintRune(122 - (r - 97))
			} else {
				z01.PrintRune(r)
			}
		}
	}
}

//compareprog

Instructions

Write a program that behaves like the Compare function from the Go package strings.

This program prints a number after comparing two string lexicographically.
Usage

student@ubuntu:~/compareprog$ go build
student@ubuntu:~/compareprog$ ./compareprog a b | cat -e
-1$
student@ubuntu:~/compareprog$ ./compareprog a a | cat -e
0$
student@ubuntu:~/compareprog$ ./compareprog b a | cat -e
1$
student@ubuntu:~/compareprog$ ./compareprog b a d | cat -e
$
student@ubuntu:~/compareprog$ ./compareprog | cat -e
$
student@ubuntu:~/compareprog$

func main() {
	defer z01.PrintRune('\n')
	neg := "-1"
	if len(os.Args) == 3 {
		a := os.Args[1]
		b := os.Args[2]
		if a > b {
			z01.PrintRune(49)
		} else if a == b {
			z01.PrintRune(48)
		} else {
			for _, r := range neg {
				z01.PrintRune(r)
			}
		}
	}
}

// doop

Instructions

Write a program that is called doop.

The program has to be used with three arguments:

    A value
    An operator
    Another value

You should use int64.

The following operators are considered valid: +, -, /, *, %.

In case of an invalid operator or overflow the programs prints 0.

In case of an invalid number of arguments the program prints nothing.

The program has to handle the modulo and division operations by 0 as shown on the output examples below.
Usage

student@ubuntu:~/[[ROOT]]/test$ go build doop.go
student@ubuntu:~/[[ROOT]]/test$ ./doop
student@ubuntu:~/[[ROOT]]/test$ ./doop 1 + 1 | cat -e
2$
student@ubuntu:~/[[ROOT]]/test$ ./doop hello + 1 | cat -e
0$
student@ubuntu:~/[[ROOT]]/test$ ./doop 1 p 1 | cat -e
0$
student@ubuntu:~/[[ROOT]]/test$ ./doop 1 / 0 | cat -e
No division by 0$
student@ubuntu:~/[[ROOT]]/test$ ./doop 1 % 0 | cat -e
No modulo by 0$
student@ubuntu:~/[[ROOT]]/test$ ./doop 9223372036854775807 + 1
0
student@ubuntu:~/[[ROOT]]/test$ ./doop -9223372036854775809 - 3
0
student@ubuntu:~/[[ROOT]]/test$ ./doop 9223372036854775807 "*" 3
0
student@ubuntu:~/[[ROOT]]/test$ ./doop 1 "*" 1
1
student@ubuntu:~/[[ROOT]]/test$ ./doop 1 "*" -1
-1
student@ubuntu:~/[[ROOT]]/test$

func main() {
	if len(os.Args) == 4 {
		var res int
		num1, f := strconv.Atoi(os.Args[1])
		num2, s := strconv.Atoi(os.Args[3])
		if f != nil || s != nil {
			fmt.Println(0)
			return
		}
		if os.Args[2] == "+" {
			res = num1 + num2
			if num1 > 0 && num2 > 0 && res < 0 {
				res = 0
			}
		}
		if os.Args[2] == "-" {
			res = num1 - num2
			if num1 < 0 && num2 < 0 && res < 0 {
				res = 0
			}
		}
		if os.Args[2] == "%" {
			if num2 == 0 {
				fmt.Println("no modulo by 0")
				return
			} else {
				res = num1 % num2
			}
		}
		if os.Args[2] == "/" {
			if num2 == 0 {
				fmt.Println("no division by 0")
				return
			} else {
				res = num1 / num2
			}
		}
		if os.Args[2] == "*" {
			res = num1 * num2
			if res/num1 != num2 {
				res = 0
			}
		}
		fmt.Println(res)
	}
}

//printchessboard

Instructions

Write a program that takes two integers as arguments and displays the chess desk, in which white cells are represented by ' ' and black cells by '#'.

    If the number of arguments is different from 2, or if the argument is not a positive number, the program displays Error followed by a newline ('\n').

Usage

student@ubuntu:~/[[ROOT]]/printchessboard$ go build
student@ubuntu:~/[[ROOT]]/printchessboard$ ./printchessboard 4 3 | cat -e
# # $
 # #$
# # $
student@ubuntu:~/[[ROOT]]/printchessboard$ ./printchessboard 7 | cat -e
Error$
student@ubuntu:~/[[ROOT]]/printchessboard$ ./printchessboard 0 0 | cat -e
Error$
student@ubuntu:~/[[ROOT]]/printchessboard$

func main() {
	err := "Error"
	if len(os.Args) == 3 {
		x, _ := strconv.Atoi(os.Args[1])
		y, _ := strconv.Atoi(os.Args[2])
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if i%2 == 0 {
					if j%2 == 0 {
						z01.PrintRune('#')
					} else {
						z01.PrintRune(' ')
					}
				} else {
					if j%2 == 0 {
						z01.PrintRune(' ')
					} else {
						z01.PrintRune('#')
					}
				}
			}
			z01.PrintRune('\n')
		}
	} else {
		for _, r := range err {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

//ispowerof2

Instructions

Write a program that determines if a given number is a power of 2.

This program must print true if the given number is a power of 2, otherwise it prints false.

    If there is more than one or no argument the program should print a newline (“\n”).

    When there is only 1 argument, it will always be a positive valid int.

Usage :

student@ubuntu:~/ispowerof2$ go build
student@ubuntu:~/ispowerof2$ ./ispowerof2 2 | cat -e
true$
student@ubuntu:~/ispowerof2$ ./ispowerof2 64 | cat -e
true$
student@ubuntu:~/ispowerof2$ ./ispowerof2 513 | cat -e
false$
student@ubuntu:~/ispowerof2$ ./ispowerof2

student@ubuntu:~/ispowerof2$ ./ispowerof2 64 1024

student@ubuntu:~/ispowerof2$

func main() {
	defer z01.PrintRune('\n')
	flag := "true"
	if len(os.Args) == 2 {
		num, _ := strconv.Atoi(os.Args[1])
		for num > 1 {
			if num%2 != 0 {
				flag = "false"
				break
			}
			num /= 2
		}
	}
	for _, r := range flag {
		z01.PrintRune(r)
	}
}

//union

Instructions

Write a program that takes two string and displays, without doubles, the characters that appear in either one of the string.

The display will be in the order that the characters will appear on the command line and will be followed by a newline ('\n').

If the number of arguments is different from 2, then the program displays newline ('\n').
Usage

student@ubuntu:~/[[ROOT]]/union$ go build
student@ubuntu:~/[[ROOT]]/union$ ./union zpadinton paqefwtdjetyiytjneytjoeyjnejeyj | cat -e
zpadintoqefwjy$
student@ubuntu:~/[[ROOT]]/union$ ./union ddf6vewg64f gtwthgdwthdwfteewhrtag6h4ffdhsd | cat -e
df6vewg4thras$
student@ubuntu:~/[[ROOT]]/union$ ./union rien "cette phrase ne cache rien" | cat -e
rienct phas$
student@ubuntu:~/[[ROOT]]/union$ ./union | cat -e
$
student@ubuntu:~/[[ROOT]]/union$ ./union rien | cat -e
$
student@ubuntu:~/[[ROOT]]/union$


func main() {
	defer z01.PrintRune('\n')
	var res string
	if len(os.Args) == 3 {
		join := os.Args[1] + os.Args[2]
		for _, r := range join {
			if isUniq(res, r) {
				res += string(r)
			}
		}
		for _, r := range res {
			z01.PrintRune(r)
		}
	}
}
func isUniq(str string, k rune) bool {
	for _, r := range str {
		if r == k {
			return false
		}
	}
	return true
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		z01.PrintRune('\n')
		return
	}
	arg1 := arg[0]
	arg2 := arg[1]
	m := map[rune]int{}
	for _, value := range arg1 + arg2 {
		if m[value] < 1 {
			m[value]++
			z01.PrintRune(rune(value))
		}
	}
	z01.PrintRune('\n')
}

//inter

Instructions

Write a program that takes two string and displays, without doubles, the characters that appear in both string, in the order they appear in the first one.

    The display will be followed by a newline ('\n').

    If the number of arguments is different from 2, the program displays a newline ('\n').

Usage

student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
padinto
student@ubuntu:~/[[ROOT]]/test$ ./test ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
df6ewg4
student@ubuntu:~/[[ROOT]]/test$


func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 3 {
		var res, fin string

		for _, b := range os.Args[1] {
			for _, r := range os.Args[2] {
				if r == b {
					res += string(r)
				}
			}
		}
		for _, r := range res {
			if isUniq(r, fin) {
				fin = fin + string(r)
			}
		}
		for _, r := range fin {
			z01.PrintRune(r)
		}
	}
}

func isUniq(r rune, str string) bool {
	for _, b := range str {
		if r == b {
			return false
		}
	}
	return true
}

func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 3 {
		str := map[rune]int{}
		for _, r := range os.Args[1] {
			for _, b := range os.Args[2] {
				if r == b && str != 1 {
					str[r]++
					fmt.Println(string(str))
				}
			}
		}
	}
}

//swapbits

##WARNING! VERY IMPORTANT!

For this exercise a function will be tested with the exam own main. However the student still needs to submit a structured program:

This means that:

    The package needs to be named package main.
    The submitted code needs one declared function main(func main()) even if empty.
    The function main declared needs to also pass the Restrictions Checker(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
    Every other rules are obviously the same than for a program.

Instructions

Write a function that takes a byte, swaps its halfs (like the example) and returns the result.
Expected function

func SwapBits(octet byte) byte {

}

Example:

1 byte

0100 | 0001
    \ /
    / \
0001 | 0100


func SwapBits(octet byte) byte {
	return octet<<4 + octet>>4
}

//twosums

nstructions

Given an array of integers, return indexes of the two numbers such that they add up to a specific target.

If there are more than one solution, return the first one.

If there are no solutions, return nil.

Expected function :

func TwoSum(nums []int, target int) []int {
}

Here is a possible program to test your function :

package main

import (
	"fmt"
)

func main() {
	case1 := []int{1, 2, 3, 4}
	out := TwoSum(case1, 5)
	fmt.Println(out)
}

And its output :

student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
[0 3]
student@ubuntu:~/[[ROOT]]/test$


func main() {
	case1 := []int{1, 2, 3, 4}
	out := TwoSum(case1, 5)
	fmt.Println(out)
}

func TwoSum(nums []int, target int) []int {
	res := []int{0, 0}
	for i, r := range nums {
		for j, k := range nums {
			if r+k == target && i != j {
				res[1] = i
				res[0] = j
			}
		}
	}
	return res
}

//repeatalpha

Instructions

Write a program called repeat_alpha that takes a string and displays it repeating each alphabetical character as many times as its alphabetical index.

The result must be followed by a newline ('\n').

'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc…

The case remains unchanged.

If the number of arguments is different from 1, the program displays a newline ('\n').
Usage

student@ubuntu:~/[[ROOT]]/repeatalpha$ go build
student@ubuntu:~/[[ROOT]]/repeatalpha$ ./repeatalpha "abc" | cat -e
abbccc
student@ubuntu:~/[[ROOT]]/repeatalpha$ ./repeatalpha "Choumi." | cat -e
CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.$
student@ubuntu:~/[[ROOT]]/repeatalpha$ ./repeatalpha "abacadaba 01!" | cat -e
abbacccaddddabba 01!$
student@ubuntu:~/[[ROOT]]/repeatalpha$ ./repeatalpha | cat -e
$
student@ubuntu:~/[[ROOT]]/repeatalpha$ ./repeatalpha "" | cat -e
$
student@ubuntu:~/[[ROOT]]/repeatalpha$


func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		str := os.Args[1]
		for _, r := range str {
			if r >= 65 && r <= 90 {
				for i := 0; i <= int(r)-65; i++ {
					z01.PrintRune(r)
				}
			}
			if r >= 97 && r <= 122 {
				for i := 0; i <= int(r)-97; i++ {
					z01.PrintRune(r)
				}
			} else {
				z01.PrintRune(r)
			}
		}
	}
}

//sortwordarrprog

Instructions

Write a function SortWordArr that sorts by ascii (in ascending order) a string array.
Expected function

func SortWordArr(array []string) {

}

Usage

Here is a possible program to test your function :

package main

import (
	"fmt"
)

func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}

And its output :

student@ubuntu:~/test$ go build
student@ubuntu:~/test$ ./test
[1 2 3 A B C a b c]
student@ubuntu:~/test$

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)
	fmt.Println(result)
}
func SortWordArr(array []string) {
	count := 0
	for range array {
		count++
	}
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

//printhex

Instructions

Write a program that takes a positive (or zero) number expressed in base 10, and displays it in base 16 (with lowercase letters) followed by a newline ('\n').

    If the number of parameters is different from 1, the program displays a newline.
    Error cases have to be handled as shown in the example below.

Usage

student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test "10"
a
student@ubuntu:~/[[ROOT]]/test$ ./test "255"
ff
student@ubuntu:~/[[ROOT]]/test$ ./test "5156454"
4eae66
student@ubuntu:~/[[ROOT]]/test$ ./test

student@ubuntu:~/[[ROOT]]/test$ ./test "123 132 1" | cat -e
0$
student@ubuntu:~/[[ROOT]]/test$


func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		base := "0123456789abcdef"
		res := ""
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			z01.PrintRune('0')
		}
		for num != 0 {
			res = string(base[num%16]) + res
			num /= 16
		}
		for _, r := range res {
			z01.PrintRune(r)
		}
	}
}

//gcd

Instructions

Write a program that takes two string representing two strictly positive integers that fit in an int.

The program displays their greatest common divisor followed by a newline ('\n').

If the number of parameters is different from 2, the program displays a newline.

All arguments tested will be positive int values.
Usage

student@ubuntu:~/[[ROOT]]/gcd$ go build
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 42 10 | cat -e
2$
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 42 12 | cat -e
6$
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 14 77 | cat -e
7$
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 17 3 | cat -e
1$
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd | cat -e
$
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 50 12 4 | cat -e
$
student@ubuntu:~/[[ROOT]]/gcd$


func main() {
	if len(os.Args) == 3 {
		res := 0
		num1, _ := strconv.Atoi(os.Args[1])
		num2, _ := strconv.Atoi(os.Args[2])
		for i := num1; i > 0; i-- {
			if num1%i == 0 && num2%i == 0 {
				res = i
				break
			}
		}
		fmt.Println(res)
	} else {
		fmt.Println()
		return
	}
}

//anagram

Instructions

Write a function that returns true if two strings are anagrams, otherwise returns false. Anagram is a string made by using the letters of another string in a different order.

Only lower case characters will be given.
Expected function

package piscine

func IsAnagram(str1, str2 string) bool {

}

Usage

Here is a possible program to test your function:

package main

import (
	piscine ".."
	"fmt"
)

func main() {
	test1 := piscine.IsAnagram("listen", "silent")
	fmt.Println(test1)

	test2 := piscine.IsAnagram("alem", "school")
	fmt.Println(test2)

	test3 := piscine.IsAnagram("neat", "a net")
	fmt.Println(test3)

	test4 := piscine.IsAnagram("anna madrigal", "a man and a girl")
	fmt.Println(test4)
}

Its output:

$> go build
$> ./main
true
false
true
true


func main() {
	test1 := IsAnagram("a", "a")
	fmt.Println(test1)

	test2 := IsAnagram("alem", "school")
	fmt.Println(test2)

	test3 := IsAnagram("abbc", "    a b bcс   ")
	fmt.Println(test3)

	test4 := IsAnagram("anna madrigal", "a man and a girl")
	fmt.Println(test4)
}

func isAnagram(str1, str2 string) bool {
	str := str1 + str2
	sum := byte(0)
	for _, v := range str {
		if v == ' ' {
			continue
		}
		sum ^= byte(v)
	}
	//fmt.Println(string(sum))
	return sum == 0
}

//balancedstring

Instructions

Balanced string is a string that has equal quantity of ‘C’ and ‘D’ characters.

Write a program that takes a string and outputs maximum amount of balanced strings without ignoring any letters. Display output with \n at the end of line.

If the number of arguments is not 1, display \n.

It will only be tested strings containing the characters ‘C’ and ‘D’.
Usage

student@ubuntu:~/[[ROOT]]/balancedstring$ go build
student@ubuntu:~/[[ROOT]]/balancedstring$ ./balancedstring "CDCCDDCDCD"
4
student@ubuntu:~/[[ROOT]]/balancedstring$ ./balancedstring "CDDDDCCCDC"
3
student@ubuntu:~/[[ROOT]]/balancedstring$ ./balancedstring "DDDDCCCC"
1
student@ubuntu:~/[[ROOT]]/balancedstring$ ./balancedstring "CDCCCDDCDD"
2

In first example “CDCCDDCDCD” can be split into “CD”, “CCDD”, “CD”, “CD”, each substring contains same number of ‘C’ and ‘D’.

Second, “CDDDDCCCDC” can be split into: “CD”, “DDDCCC”, “DC”.

“DDDDCCCC” can be split into “DDDDCCCC”.

“CDCCCDDCDD” can be split into: “CD”, “CCCDDCDD”, since each substring contains an equal number of ‘C’ and ‘D’

func main() {
	if len(os.Args) == 2 {
		var count, res int
		for _, r := range os.Args[1] {
			if r == 'C' {
				count++
			} else {
				count--
			}
			if count == 0 {
				res++
			}
		}
		fmt.Println(res)
	} else {
		fmt.Println()
		return
	}
}

// lcm

Instructions

Write a function, lcm, that returns least common multiple.

It will be tested with positive int values and 0.
Expected function

func Lcm(first, second int) int {

}

Usage

Here is a possible program to test your function :

package main

func main() {
	fmt.Println(Lcm(2, 7))
	fmt.Println(Lcm(0, 4))
}

Output

student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
14
0
student@ubuntu:~/[[ROOT]]/test$

func main() {
	fmt.Println(Lcm(7, 34))
}
func Lcm(first, second int) int {
	res := 0
	for i := second; i <= first*second; i++ {
		if i%first == 0 && i%second == 0 {
			res = i
			break
		}
	}
	return res
}

//nauuo

Instructions

There was a vote. There are people who voted positively, negatively, and randomly. Figure out if the final answer depends on random people or not. If it does, return ‘?’, otherwise the result must be either ‘+’, ‘-‘, or ‘0’ Previous characters stand for outcome of the vote: positive/negative/draw. Input is always positive.

Write a function, Nauuo, that returns final result of voting.
Expected function

func Nauuo(plus, minus, rand int) string {

}

Usage

Here is a possible program to test your function :

package main

import (
	"fmt"
	piscine ".."
)

func main() {
	fmt.Println(piscine.Nauuo(50, 43, 20))
	fmt.Println(piscine.Nauuo(13, 13, 0))
	fmt.Println(piscine.Nauuo(10, 9, 0))
	fmt.Println(piscine.Nauuo(5, 9, 2))
}

And its output :

student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
?
0
+
-
student@ubuntu:~/[[ROOT]]/test$


func main() {
	fmt.Println(Nauuo(50, 43, 20))
	fmt.Println(Nauuo(13, 13, 0))
	fmt.Println(Nauuo(10, 9, 0))
	fmt.Println(Nauuo(5, 9, 2))
}
func Nauuo(plus, minus, rand int) string {
	var res string
	if plus == minus && rand == 0 {
		res = "0"
	} else if plus > (minus + rand) {
		res = "+"
	} else if (plus + rand) < minus {
		res = "-"
	} else {
		res = "?"
	}
	return res
}

//uniqueoccur

Instructions

Write a program that outputs true if the number of occurrences of each character is unique, otherwise false. \n should be at the of line.

If number of arguments is not 1 output \n.

Only lower case characters will be given.
Usage

student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./main "abbaac"
true
student@ubuntu:~/[[ROOT]]/test$ ./main "ab"
false
student@ubuntu:~/[[ROOT]]/test$ ./main "abcacccazb"
true

In first example, character ‘a’ has 3 occurrences, ‘b’ has 2 and ‘c’ has 1. No two characters have the same number of occurrences.

func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		str := os.Args[1]
		m := map[rune]int{}
		res := ""
		for _, r := range str {
			m[r]++
		}
		fmt.Println(m)
		for key, count := range m {
			if isUniq(m, key, count) {
				toPrint("true")
				res = "true"
			} else {
				toPrint("false")
				return
			}
		}
		toPrint(res)
	}
}
func toPrint(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
}
func isUniq(m map[rune]int, key rune, count int) bool {
	for s, num := range m {
		if s != key {
			if num == count {
				return false
			}
		}
	}
	return true
}

//hiddenp

Instructions

Write a program named hiddenp that takes two string and that, if the first string is hidden in the second one, displays 1 followed by a newline ('\n'), otherwise it displays 0 followed by a newline.

Let s1 and s2 be string. It is considered that s1 is hidden in s2 if it is possible to find each character from s1 in s2, in the same order as they appear in s1.

If s1 is an empty string it is considered hidden in any string.

If the number of parameters is different from 2, the program displays a newline.
Usage

student@ubuntu:~/[[ROOT]]/hiddenp$ go build
student@ubuntu:~/[[ROOT]]/hiddenp$ ./hiddenp "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
1$
student@ubuntu:~/[[ROOT]]/hiddenp$ ./hiddenp "abc" "2altrb53c.sse" | cat -e
1$
student@ubuntu:~/[[ROOT]]/hiddenp$ ./hiddenp "abc" "btarc" | cat -e
0$
student@ubuntu:~/[[ROOT]]/hiddenp$ ./hiddenp "DD" "DABC" | cat -e
0$
student@ubuntu:~/[[ROOT]]/hiddenp$ ./hiddenp | cat -e
$
student@ubuntu:~/[[ROOT]]/hiddenp$

func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 3 {
		str1 := os.Args[1]
		str2 := os.Args[2]
		i := 0
		for j := range str2 {
			if str1[i] == str2[j] {
				i++
				if i == len(str1) {
					z01.PrintRune('1')
					return
				}
			}
		}
		z01.PrintRune('0')
	}
}

// rostring

Instructions

Write a program that takes a string and displays this string after rotating it one word to the left.

Thus, the first word becomes the last, and others stay in the same order.

A word is a sequence of alphanumerical characters.

Words will be separated by only one space in the output.

If the number of arguments is different from 1, the program displays a newline.
Usage

student@ubuntu:~/[[ROOT]]/rostring$ go build
student@ubuntu:~/[[ROOT]]/rostring$ ./rostring "abc   " | cat -e
abc$
student@ubuntu:~/[[ROOT]]/rostring$ ./rostring "Let there     be light"
there be light Let
student@ubuntu:~/[[ROOT]]/rostring$ ./rostring "     AkjhZ zLKIJz , 23y"
zLKIJz , 23y AkjhZ
student@ubuntu:~/[[ROOT]]/rostring$ ./rostring | cat -e
$
student@ubuntu:~/[[ROOT]]/rostring$


func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		res := swap(splitWS(os.Args[1]))
		res1 := ""
		for i, r := range res {
			if i != len(res)-1 {
				res1 = res1 + r + " "
			} else {
				res1 = res1 + r
			}
		}
		for _, b := range res1 {
			z01.PrintRune(b)
		}
	}
}

func swap(arg []string) []string {
	counter := 0
	for i := range arg {
		counter++
		if counter != len(arg) {
			arg[i], arg[counter] = arg[counter], arg[i]
		}
	}
	return arg
}
func splitWS(str string) []string {
	var s string
	count := 0
	n := 0
	str = str + " "
	for i := range str {
		if str[i] == ' ' {
			s = str[n:i]
			n = i + 1
			if s != "" {
				count++
			}
		}
	}
	n = 0
	arr := make([]string, count)
	j := 0
	for i := range str {
		if str[i] == ' ' {
			s = str[n:i]
			n = i + 1
			if s != "" {
				arr[j] = s
				j++
			}
		}
	}
	return arr
}

//splitprog

Instructions

Write a function that separates the words of a string and puts them in a string array.

The separators are the characters of the charset string given in parameter.
Expected function

func Split(str, charset string) []string {

}

Usage

Here is a possible program to test your function :

package main

import (
	"fmt"
	piscine ".."
)

func main() {
	str := "HelloHAhowHAareHAyou?"
	fmt.Println(piscine.Split(str, "HA"))
}

And its output :

student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
[Hello how are you?]
student@ubuntu:~/[[ROOT]]/test$

func main() {
	str := "HelloHAhowHAareHAyou?"
	fmt.Println(Split(str, "HA"))
}
func Split(str, charset string) []string {
	if charset == "" {
		ar := make([]string, len(str))
		i := 0
		for _, v := range str {
			ar[i] = string(v)
			i++
		}
		return ar
	}
	nw := 0
	a := 0
	for i := 0; i <= len(str)-len(charset); i++ {
		if str[i:i+len(charset)] == charset {
			a++
			i = i + len(charset)
		}
	}
	ar := make([]string, a+1)
	j := 0
	for i := 0; i <= len(str)-len(charset); i++ {
		if str[i:i+len(charset)] == charset {
			ar[j] = str[nw:i]
			j++
			nw = i + len(charset)
			i = i + len(charset)
		}
	}
	ar[j] = str[nw:]
	return ar
}

//revwstr

Instructions

Write a program that takes a string as a parameter, and prints its words in reverse.

    A word is a sequence of alphanumerical characters.

    If the number of parameters is different from 1, the program will display newline ('\n').

    In the parameters that are going to be tested, there will not be any extra spaces. (meaning that there will not be additional spaces at the beginning or at the end of the string and that words will always be separated by exactly one space).

Usage

student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test "the time of contempt precedes that of indifference"
indifference of that precedes contempt of time the
student@ubuntu:~/[[ROOT]]/test$ ./test "abcdefghijklm"
abcdefghijklm
student@ubuntu:~/[[ROOT]]/test$ ./test "he stared at the mountain"
mountain the at stared he
student@ubuntu:~/[[ROOT]]/test$ ./test "" | cat-e
$
student@ubuntu:~/[[ROOT]]/test$ 


func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		var res string
		str := os.Args[1]
		for i := len(str) - 1; i != 0; i-- {
			if str[i] == ' ' {
				res += str[i+1:] + " "
				str = str[:i]
			}
		}
		res += str
		for _, r := range res {
			z01.PrintRune(r)
		}
	}
}

//priorprime

Instructions

You are given an integer. Your function must return sum of all prime numbers prior to the number exclusively. The number is not included.
Expected function and structure

package main

func Priorprime(x int) int {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Priorprime(14))
}

Its output:

student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./priorprime
41
student@ubuntu:~/[[ROOT]]/test$

func main() {
	fmt.Println(Prioprime(14))
}
func Prioprime(x int) int {
	res := 0
	for i := 2; i < x; i++ {
		if Prime(i) {
			fmt.Println(i)
			res += i
		}
	}
	return res
}
func Prime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

//fprime

Instructions

Write a program that takes a positive int and displays its prime factors, followed by a newline ('\n').

    Factors must be displayed in ascending order and separated by *.

    If the number of parameters is different from 1, the program displays a newline.

    The input, when there is one, will always be valid.

Usage

student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test 225225
3*3*5*5*7*11*13
student@ubuntu:~/[[ROOT]]/test$ ./test 8333325
3*3*5*5*7*11*13*37
student@ubuntu:~/[[ROOT]]/test$ ./test 9539
9539
student@ubuntu:~/[[ROOT]]/test$ ./test 804577
804577
student@ubuntu:~/[[ROOT]]/test$ ./test 42
2*3*7
student@ubuntu:~/[[ROOT]]/test$ ./test a

student@ubuntu:~/[[ROOT]]/test$ ./test 0

student@ubuntu:~/[[ROOT]]/test$ ./test 1

student@ubuntu:~/[[ROOT]]/test$


func main() {
    arg := os.Args[1]
    divider := 2
    num, _ := strconv.Atoi(arg)
    for divider <= num {
        if num%divider == 0 {
            fmt.Print(divider)
            if divider == num {
                fmt.Println("")
                return
            }
            fmt.Print("*")
            num /= divider
            divider = 1
        }
        divider++
    }
}
