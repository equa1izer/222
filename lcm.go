package main

func Lcm(first, second int) int {
	x := 0
	if first == 0 || second == 0 {
		return 0
	}
	for i := 1; i <= second; i++ {
		x = i * first
		if x%second == 0 {
			return x
		}
	}
	return first * second
}
