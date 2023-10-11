package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

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
