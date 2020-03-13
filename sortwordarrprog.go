package main

func main() {

}

func SortWordArr(array []string) {
	l := 0
	for range array {
		l++
	}
	for i := 0; i < l-1; i = i {
		z := array[i]
		if array[i] > array[i+1] {
			array[i] = array[i+1]
			array[i+1] = z
			i = 0
		} else {
			i++
		}
	}

}
