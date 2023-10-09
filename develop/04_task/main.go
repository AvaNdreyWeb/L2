package main

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
