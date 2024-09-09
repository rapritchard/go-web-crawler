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

	rawBaseURL := args[0]
	const maxConcurrency = 3
	cfg, err := configure(rawBaseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}
	fmt.Printf("starting crawl of: %s...\n", args[0])

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for url, visits := range cfg.pages {
		fmt.Printf("%s ->  %d\n", url, visits)
	}
}
