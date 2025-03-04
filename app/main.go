package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Something went wrong")
	}
	command = command[:len(command)-1]

	// Split the command into words.
	args := strings.Fields(command)

	switch args[0] {
	case "exit":
		// if the status_code is not provided, exit with the code 0.
		if len(args) == 1 {
			os.Exit(0)
		}

		status_code, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Invalid argument %s", args[1])
			os.Exit(0)
		}
		os.Exit(status_code)

	case "echo":
		fmt.Println(strings.Join(args[1:], " "))

	case "type":
		// Verify if the command name is provided.
		if len(args) == 1 {
			fmt.Println("Missing command name")
		} else if BuiltinsMap[args[1]] {
			fmt.Printf("%s is a shell builtin\n", args[1])
		} else {
			fmt.Printf("%s: not found\n", args[1])
		}

	default:
		fmt.Println(command + ": command not found")
	}

	// Create a REPL.
	main()
}
