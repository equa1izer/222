package main

import (
    "os"

    "github.com/01-edu/z01"
)

func main() {
    if len(os.Args[1:]) != 1 {
        z01.PrintRune('\n')
        return
    } else {
        s := os.Args[1]
        res := ""
        for i := len(s) - 1; i >= 0; i-- {
            if s[i] == ' ' {
                res += s[i+1:] + " "
                s = s[:i]
            }
        }
        res += s
        for _, v := range res {
            z01.PrintRune(v)
        }
        z01.PrintRune('\n')
    }
}
