package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gocoreutils <command> [arguments]")
		os.Exit(1)
	}

	command := os.Args[1] // first argument is the command

	switch command {
	case "cat":
		Cat(os.Args[2:])
	case "ls":
		Ls(os.Args[2:])
	case "echo":
		Echo(os.Args[2:])
	case "pwd":
		Pwd(os.Args[2:])
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}
