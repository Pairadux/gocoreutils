package main

import (
	"fmt"
	"os"
)

func Ls(args []string) {
	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}
