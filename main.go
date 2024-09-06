package main

import (
	"fmt"
	"os"
)

func main() {
	// args without the path to the program
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[0]
	fmt.Printf("starting crawl of: %v\n", baseURL)
	html, err := getHTML(baseURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(html)
}
