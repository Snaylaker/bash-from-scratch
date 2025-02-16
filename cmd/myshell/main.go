package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")
	allowedCmd := []string{"ls", "cd"}
	// Wait for user input
	for {
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		cmd = cmd[:len(cmd)-1]

		if !slices.Contains(allowedCmd, cmd) {
			fmt.Printf("%s: command not found\n", cmd)
			continue
		}
	}
}
