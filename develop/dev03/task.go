package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type AppArgs struct {
	inputFile  string
	outputFile string
	r          bool
	u          bool
	n          bool
	k          int
}

func ParseArgs() *AppArgs {
	if len(os.Args) < 2 {
		return nil
	}
	O := flag.String(
		"o",
		"sorted.txt",
		"Output file",
	)
	R := flag.Bool("r", false, "Reverse sorting")
	U := flag.Bool("u", false, "Unique strings only")
	N := flag.Bool("n", false, "Numerical sorting")
	K := flag.Int("k", 0, "Column to sort")
	flag.Parse()
	I := flag.Arg(0)
	if *K < 0 {
		*K = 0
	}
	return &AppArgs{
		inputFile:  I,
		outputFile: *O,
		r:          *R,
		u:          *U,
		n:          *N,
		k:          *K,
	}
}

func main() {
	args := ParseArgs()
	if args == nil {
		fmt.Println("sort [-o <filename> -k <column> -r -u -n] <filename>")
		os.Exit(1)
	}

	data, err := ReadFile(args.inputFile)
	if err != nil {
		panic(err)
	}

	sorted := SortData(data, args.n, args.k)
	err = WriteDataToFile(args.outputFile, sorted, args.r, args.u)
	if err != nil {
		panic(err)
	}
}

func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	data := []string{}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return []string{}, err
		}
		data = append(data, line)
	}
	return data, nil
}

func WriteDataToFile(filename string, data []string, desc, uniq bool) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var prev string
	n := len(data)
	for i := 0; i < n; i++ {
		var line string
		if desc {
			line = data[n-i-1]
		} else {
			line = data[i]
		}
		if (uniq && line != prev) || (!uniq) {
			fmt.Fprintln(file, line)
		}
		prev = line
	}
	return nil
}

func SortData(data []string, num bool, col int) []string {
	tmp := [][]string{}
	for _, line := range data {
		tmp = append(tmp, strings.Fields(line))
	}
	sort.Slice(tmp, func(i, j int) bool {
		if num {
			x, _ := strconv.Atoi(tmp[i][col])
			y, _ := strconv.Atoi(tmp[j][col])
			if x < y {
				return true
			}
		} else {
			if tmp[i][col] < tmp[j][col] {
				return true
			}
		}
		return false
	})

	res := []string{}
	for _, fields := range tmp {
		res = append(res, strings.Join(fields, " "))
	}

	return res
}
