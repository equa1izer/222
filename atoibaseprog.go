package main

import "fmt"

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	} else if power%2 == 0 {
		return RecursivePower(nb, power/2) * RecursivePower(nb, power/2)
	} else {
		return nb * RecursivePower(nb, power/2) * RecursivePower(nb, power/2)
	}
}

func IsUnique(base string) bool {
	length_base1 := 0
	for range base {
		length_base1++
	}
	for index, letter := range base {
		for i := index + 1; i < length_base1; i++ {
			if letter == '-' || letter == '+' || letter == rune(base[i]) {
				return false
			}
		}
	}
	return true
}

func Index(s string, toFind string) int {
	j := []rune(s)
	j_count := 0
	l := []rune(toFind)
	l_count := 0
	for range j {
		j_count++
	}
	for range l {
		l_count++
	}
	n := j_count
	k := l_count
	for i := 0; i <= n-k; i++ {
		if toFind == s[i:i+k] {
			return (i)
		}
	}
	return -1
}

func AtoiBase(s string, base string) int {
	indx := 0
	if IsUnique(base) == false {
		indx = 1
	}
	if indx == 1 || len(base) < 2 {
		return 0
	} else {
		result := 0
		for i, res := range s {
			ind := Index(base, string(res))
			result = result + ind*RecursivePower(len(base), len(s)-1-i)
		}
		return result
	}
}
