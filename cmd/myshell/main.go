package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type BuiltinCommand string

const (
	Echo BuiltinCommand = "echo"
	Exit BuiltinCommand = "exit"
	Type BuiltinCommand = "type"
)

var builtins = []BuiltinCommand{Echo, Exit, Type}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		if !scanner.Scan() {
			break
		}
		cmd := strings.TrimSpace(scanner.Text())
		handleCommand(cmd)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
}

func handleCommand(cmd string) {
	if strings.HasPrefix(cmd, string(Echo)+" ") {
		fmt.Println(strings.TrimPrefix(cmd, string(Echo)+" "))
		return
	}

	if strings.HasPrefix(cmd, string(Type)+" ") {
		cmdToCheck := strings.TrimPrefix(cmd, string(Type)+" ")
		if slices.Contains(builtins, BuiltinCommand(cmdToCheck)) {
			fmt.Printf("%s is a shell builtin\n", cmdToCheck)
		} else {
			fmt.Printf("%s: not found\n", cmdToCheck)
		}
		return
	}

	if cmd == "exit 0" {
		os.Exit(0)
	}

	fmt.Printf("%s: command not found\n", cmd)
}
