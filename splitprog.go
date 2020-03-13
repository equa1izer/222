package main

import "fmt"

func main() {
	str := "HelloHAhowHAareHAyou?"
	fmt.Println(Split(str, "HA"))
}

func Split(str, charset string) []string {
	end := str + charset
	nbword := 0
	for i := 0; i <= len(string(end))-len(string(charset)); i++ {
		if end[i:i+len(string(charset))] == charset {
			nbword++
		}
	}

	arr := make([]string, nbword)
	index := 0
	arrIndex := 0
	for i := index; i <= len(string(end))-len(string(charset)); i++ {
		if end[i:i+len(string(charset))] == charset {
			arr[arrIndex] = end[index:i]
			index = i + len(string(charset))
			arrIndex++
		}
	}
	return arr
}
