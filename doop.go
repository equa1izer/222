package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 4 {
		v1, err := strconv.Atoi(os.Args[1])
		x := 0
		if err != nil {
			fmt.Println("0")
			return
		}
		v2, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("0")
			return
		}
		if v1 < 0 && v2 < 0 && v1+v2 > 0 {
			fmt.Println("0")
			return
		}
		if v1 > 0 && v2 > 0 && v1+v2 < 0 {
			fmt.Println("0")
			return
		}
		if os.Args[2] == "+" {
			x = v1 + v2

		} else if os.Args[2] == "-" {
			x = v1 - v2

		} else if os.Args[2] == "*" {
			x = v1 * v2

		} else if os.Args[2] == "/" {
			if v2 == 0 {
				fmt.Println("No division by 0")
				return
			} else {
				x = v1 / v2
			}

		} else if os.Args[2] == "%" {
			if v2 == 0 {
				fmt.Println("No modulo by 0")
				return
			}
			x = v1 % v2
		}
		fmt.Println(x)
	}

}
