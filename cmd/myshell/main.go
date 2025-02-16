package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	allowedCmd := []string{"ls", "cd", "echo"}
	// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if strings.HasPrefix(cmd, "echo ") {
			result := strings.TrimPrefix(cmd, "echo ")
			fmt.Print(result + "\r")
			continue
		}

		cmd = cmd[:len(cmd)-1]

		if cmd == "exit 0" {
			os.Exit(0)
		}

		if !slices.Contains(allowedCmd, cmd) {
			fmt.Printf("%s: command not found\n", cmd)
			continue
		}
		continue
	}
}
