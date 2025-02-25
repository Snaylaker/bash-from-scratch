package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type BuiltinCommand = string

const (
	Echo BuiltinCommand = "echo"
	Exit BuiltinCommand = "exit"
	Type BuiltinCommand = "type"
)

var builtins = []BuiltinCommand{Echo, Exit, Type}

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if strings.HasPrefix(cmd, Echo) {
			result := strings.TrimPrefix(cmd, "echo ")
			fmt.Print(result + "\r")
		}

		cmd = cmd[:len(cmd)-1]

		if strings.HasPrefix(cmd, "type ") {
			cmdToCheck := strings.TrimPrefix(cmd, "type ")
			if cmd == Echo {
				fmt.Printf("%s is a shell builtin\n", cmdToCheck)
			} else {
				fmt.Printf("%s: not found\n", cmdToCheck)
			}
		}

		if cmd == "exit 0" {
			os.Exit(0)
		}

		fmt.Printf("%s: command not found\n", cmd)
		continue
	}
}
