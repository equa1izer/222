package main

func BalancedString(str string) int {
	counter := 0
	balanceCounter := 0

	for i := 0; i < len(str); i++ {
		c := str[i]
		if c == 'C' {
			balanceCounter++
		} else if c == 'D' {
			balanceCounter--
		}

		if balanceCounter == 0 {
			counter++
		}
	}

	return counter
}
