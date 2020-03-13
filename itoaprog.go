package main

func main() {

}

func Itoa(n int) string {

	var result string
	flag := false

	if n < 0 {
		n = -n
		flag = true

	}

	for n > 0 {
		d := (n % 10) + 48
		result = string(d) + result
		n = (n / 10)
	}

	if flag {
		result = "-" + result
	}
	return result

}
