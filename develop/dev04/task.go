package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	data := []string{
		"пятка",
		"тяпка",
		"пятак",
		"листок",
		"Листок",
		"слиток",
		"столик",
		"Сталик",
		"сталик",
	}
	ans := findAnagrams(data)
	fmt.Println(ans)
}

func findAnagrams(strs []string) *map[string][]string {
	res := make(map[string][]string)
	for _, word := range strs {
		low := strings.ToLower(word)
		new := true
		for key := range res {
			if !isSliceContain(res[key], low) && isAnagram(key, low) {
				res[key] = append(res[key], low)
				new = false
			}
		}
		if new {
			res[low] = []string{}
			res[low] = append(res[low], low)
		}
	}
	toDel := []string{}
	for key, value := range res {
		if len(value) == 1 {
			toDel = append(toDel, key)
		}
	}
	for _, key := range toDel {
		delete(res, key)
	}
	for key := range res {
		sort.Strings(res[key])
	}

	return &res
}

func isAnagram(str1, str2 string) bool {
	check := make(map[rune]int)
	for _, r := range str1 {
		check[r] += 1
	}
	for _, r := range str2 {
		check[r] -= 1
	}
	for _, v := range check {
		if v != 0 {
			return false
		}
	}
	return true
}

func isSliceContain(strs []string, str string) bool {
	for _, v := range strs {
		if v == str {
			return true
		}
	}
	return false
}
