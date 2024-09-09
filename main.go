package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// args without the path to the program
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: go-web-crawler <baseURL> <maxConcurrency> <maxPages>")
		os.Exit(1)
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Error - maxConcurrency: %v", err)
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Error - maxPages: %v", err)
		os.Exit(1)
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}
	fmt.Printf("starting crawl of: %s...\n", args[0])

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, rawBaseURL)
}
