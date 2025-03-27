package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
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
		} else if path, err := exec.LookPath(args[1]); err == nil {
			fmt.Printf("%s is %s\n", args[1], path)
		} else {
			fmt.Printf("%s: not found\n", args[1])
		}

	case "pwd":
		if dir, err := os.Getwd(); err == nil {
			fmt.Println(dir)
		} else {
			fmt.Println("Something went wrong!")
		}

	case "cd":
		if len(args) < 2 {
			fmt.Println("Please specify target path")
		} else {
			var path string
			if strings.HasPrefix(args[1], "~") {
				path, _ = os.UserHomeDir()
			} else {
				path = args[1]
			}
			err := os.Chdir(path)
			if err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", args[1])
			}
		}

	default:
		executableProcess := false

		// Check if the first argument passed is an executable file.
		if len(args) >= 1 {
			paths := strings.Split(os.Getenv("PATH"), ":")

			// Iterate through all the occurences of PATH and check if the program is found.
			for i := range len(paths) {
				fullExecutablePath := paths[i] + "/" + args[0]

				// Logic to determine if the file is already present.
				if fileInfo, err := os.Stat(fullExecutablePath); err == nil {
					// Check if file is executable
					if fileInfo.Mode()&0111 != 0 {
						cmd := exec.Command(args[0], args[1:]...)
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr
						cmd.Stdin = os.Stdin
						if err := cmd.Run(); err != nil {
							log.Fatal(err)
						}

						executableProcess = true
						break
					}
				}
			}
		}

		if !executableProcess {
			fmt.Println(command + ": command not found")
		}
	}

	// Create a REPL.
	main()
}
