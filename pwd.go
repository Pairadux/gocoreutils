package main

import (
	"fmt"
	"os"
)

func Pwd(args []string) {
	if len(args) > 0 {
		fmt.Println("pwd: too many arguments")
		os.Exit(1)
	}
	fmt.Println(os.Getwd())
}
