package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(UnpackString("a4bc2d5e"))
	fmt.Println(UnpackString("abcd"))
	fmt.Println(UnpackString("45"))
	fmt.Println(UnpackString(""))
}

func UnpackString(str string) string {
	var last rune = -1
	unpacked := ""

	for _, r := range str {
		num, err := strconv.Atoi(string(r))
		if err != nil {
			last = r
			unpacked += string(last)
		} else {
			if last == -1 {
				return ""
			}
			for i := 0; i < num-1; i++ {
				unpacked += string(last)
			}
			last = -1
		}
	}
	return unpacked
}
