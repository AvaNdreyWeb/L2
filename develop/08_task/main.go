package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"os/exec"
	"strings"
)

var shutdown = errors.New("user exit")

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	io.WriteString(os.Stdout, ">>> ")
	for scanner.Scan() {
		req := scanner.Text()
		res, err := CommandHandler(req)
		if err == shutdown {
			io.WriteString(os.Stdout, res[0])
			io.WriteString(os.Stdout, "\n")
			os.Exit(0)
		}
		for _, line := range res {
			io.WriteString(os.Stdout, line)
			io.WriteString(os.Stdout, "\n")
		}
		io.WriteString(os.Stdout, ">>> ")
	}
}

func CommandHandler(req string) ([]string, error) {
	args := strings.Fields(req)
	main := args[0]
	res := []string{}
	switch main {
	case "\\quit":
		res = append(res, "Завершение работы")
		return res, shutdown
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return res, err
		}
		res = append(res, dir)
	case "cd":
		if len(args) > 1 {
			dir := args[1]
			err := os.Chdir(dir)
			if err != nil {
				res = append(res, "no such directory: "+dir)
			}
		} else {
			res = append(res, "help: cd <args>")
		}
	case "echo":
		if len(args) > 1 {
			echo := args[1]
			res = append(res, echo)
		} else {
			res = append(res, "help: echo <args>")
		}
	case "kill":
		if len(args) > 1 {
			pid := args[1]
			cmd := exec.Command("kill", pid)
			out, err := cmd.Output()
			if err != nil {
				return res, err
			}
			res = append(res, string(out))
		} else {
			res = append(res, "help: kill <pid>")
		}
	case "ps":
		cmd := exec.Command("ps", "aux")
		out, err := cmd.Output()
		if err != nil {
			return res, err
		}
		res = append(res, string(out))
	default:
		if len(args) < 2 {
			res = append(res, "help: cmd <args>")
		} else {
			cmd := exec.Command(args[0], args[1:]...)
			out, err := cmd.Output()
			if err != nil {
				return res, err
			}
			res = append(res, string(out))
		}
	}
	return res, nil
}
