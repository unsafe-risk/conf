package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: clean <dir or file>")
		os.Exit(1)
	}

	stat, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	if stat.IsDir() {
		err = os.RemoveAll(os.Args[1])
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		os.Mkdir(os.Args[1], stat.Mode())
	} else {
		err = os.Remove(os.Args[1])
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
	}
}
