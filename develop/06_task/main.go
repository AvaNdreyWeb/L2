package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type inputFlags struct {
	fields        []int
	delimiter     string
	filename      string
	onlyWithDelim bool
}

func main() {
	params, err := parseArgs()
	if err != nil {
		fmt.Println("Invalid arg!", err)
		return
	}

	data, err := readFromFile(params.filename)
	if err != nil {
		log.Println(err)
	}

	cut(data, params, os.Stdout)
}

func parseArgs() (inputFlags, error) {
	var (
		fields    = flag.String("f", "", "Choose fields")
		delimiter = flag.String("d", "", "Selected delimiter")
		sep       = flag.Bool("s", false, "Separator")
	)
	flag.Parse()

	if len(*delimiter) != 1 {
		return inputFlags{}, fmt.Errorf("Delim should be with size pattern")
	}
	if len(flag.Args()) != 1 {
		return inputFlags{}, fmt.Errorf("Filename should be provided")
	}

	cleanFields, err := parseFields(*fields)
	if err != nil {
		log.Fatal(err)
	}

	params := inputFlags{
		fields:        cleanFields,
		delimiter:     *delimiter,
		onlyWithDelim: *sep,
		filename:      flag.Args()[0],
	}

	return params, nil
}

func readFromFile(inputFile string) ([]string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	line := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return line, nil
}

func parseFields(fields string) ([]int, error) {
	var bucket = make([]int, 0)
	if strings.ContainsAny(fields, ",") {
		bucket, err := strToInt(",", fields)
		return bucket, err

	} else if strings.ContainsAny(fields, "-") {
		tmpBucket, err := strToInt("-", fields)
		for i := tmpBucket[0]; i < tmpBucket[1]+1; i++ {
			bucket = append(bucket, i)
		}
		return bucket, err

	} else {
		bucket, err := strToInt(" ", fields)
		return bucket, err
	}
}

func strToInt(chr, fields string) ([]int, error) {
	var bucket = make([]int, 0)
	for _, val := range strings.Split(fields, chr) {
		val, err := strconv.Atoi(val)
		if err != nil {
			return []int{}, err
		}
		bucket = append(bucket, val)
	}
	return bucket, nil
}

func cut(data []string, params inputFlags, out io.Writer) {
	for i := 0; i < len(data); i++ {
		delimExist := strings.Contains(data[i], params.delimiter)
		if params.onlyWithDelim && !delimExist {
			continue
		}
		if !delimExist && !params.onlyWithDelim {
			fmt.Fprintf(out, "%s\n", data[i])
			continue
		}
		column := strings.Split(data[i], params.delimiter)
		for _, num := range params.fields {
			if num > len(column) {
				break
			}

			fmt.Fprintf(out, "%s", column[num-1])

			if num != params.fields[len(params.fields)-1] && num < len(column) {
				fmt.Fprintf(out, "%s", params.delimiter)
			}
		}
		fmt.Fprint(out, "\n")
	}
}
