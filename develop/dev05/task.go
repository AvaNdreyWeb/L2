package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type inputFlags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	inversion  bool
	lineNum    bool
	pattern    string
	fileName   string
}

func main() {
	params, err := parsArgs()
	if err != nil {
		fmt.Println("Invalid arg!", err)
	}

	data, err := readFromFile(params.fileName)
	if err != nil {
		log.Println(err)
	}

	grep(data, params, os.Stdout)
}

func parsArgs() (inputFlags, error) {
	var (
		after      = flag.Int("A", 0, "+N string after context")
		before     = flag.Int("B", 0, "+N string before context")
		context    = flag.Int("C", 0, "+/- N string around context")
		count      = flag.Bool("c", false, "Count of string of context")
		ignoreCase = flag.Bool("i", false, "Ignore context case")
		inversion  = flag.Bool("v", false, "Inverse context")
		lineNum    = flag.Bool("n", false, "Num context line")
	)

	flag.Parse()

	flags := inputFlags{
		after:      *after,
		before:     *before,
		count:      *count,
		ignoreCase: *ignoreCase,
		inversion:  *inversion,
		lineNum:    *lineNum,
		pattern:    "c",
		fileName:   "test.txt",
	}

	if *context != 0 {
		flags.after = *context
		flags.before = *context
	}

	return flags, nil
}

func readFromFile(inputFile string) (map[int]string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return map[int]string{}, err
	}
	defer file.Close()

	var (
		table = make(map[int]string, 0)
		i     = 0
	)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		table[i] = scanner.Text()
		i++
	}

	if err := scanner.Err(); err != nil {
		return map[int]string{}, err
	}

	return table, nil
}

func grep(data map[int]string, params inputFlags, out io.Writer) {
	var (
		keys    = make([]int, 0)
		allKeys = make([]int, 0)
		res     = make([]int, 0)
		ln      = len(data)
	)

	for key, val := range data {
		allKeys = append(allKeys, key)
		if params.ignoreCase {
			val = strings.ToLower(val)
		}
		if strings.Contains(val, params.pattern) {
			keys = append(keys, key)
		}
	}

	sort.Ints(keys)

	if params.inversion {
		keys = invert(keys, allKeys)
	}

	if params.before > 0 {
		res = append(res, before(keys, params.before, ln)...)
	}

	if params.after > 0 {
		res = append(res, after(keys, params.after, ln)...)
	}

	if params.before == 0 && params.after == 0 && !params.count {
		for _, val := range keys {
			fmt.Fprintf(out, "%s\n", data[val])
		}
	}

	res = set(res)
	if !params.count {
		for k, val := range res {
			if params.lineNum {
				fmt.Fprintf(out, "%d:%s\n", k+1, data[val])
			} else {
				fmt.Fprintf(out, "%s\n", data[val])
			}
		}
	} else {
		fmt.Fprintf(out, "%d\n", len(keys))
	}
}

func set(data []int) []int {
	set := make(map[int]bool)
	var setToSlice []int
	for _, val := range data {
		set[val] = true
	}
	for i := range set {
		setToSlice = append(setToSlice, i)
	}
	sort.Ints(setToSlice)
	return setToSlice
}

func after(keys []int, aft, ln int) []int {
	var res []int
	for _, v := range keys {
		after := v + aft
		for i := v; i <= after; i++ {
			if ln > i {
				res = append(res, i)
			}
		}
	}
	return set(res)
}

func before(keys []int, aft, ln int) []int {
	var res []int
	for _, v := range keys {
		after := v - aft
		if after < 0 {
			after = 0
		}
		for i := v; i >= after; i-- {
			if ln > i {
				res = append(res, i)
			}
		}
	}
	return set(res)
}

func invert(arr []int, allKeys []int) []int {
	sort.Ints(allKeys)
	for i, val := range arr {
		allKeys = append(allKeys[:val-i], allKeys[val-i+1:]...)
	}
	return allKeys
}
