package main

import (
	"fmt"
	"os"
	"io"
)

func Cat(args []string) {
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for _, file := range args {
		f, err := os.Open(file)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		io.Copy(os.Stdout, f)
		f.Close()
	}
}
