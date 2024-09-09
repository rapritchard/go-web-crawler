package main

import (
	"fmt"
	"sort"
)

type Page struct {
	URL    string
	Visits int
}

func printReport(pages map[string]int, baseURL string) {

	fmt.Println("=============================")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	sortedPages := sortPages(pages)
	for _, p := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", p.Visits, p.URL)
	}
}

func sortPages(pages map[string]int) []Page {
	pageSlice := []Page{}
	for url, visits := range pages {
		pageSlice = append(pageSlice, Page{URL: url, Visits: visits})
	}

	sort.Slice(pageSlice, func(i, j int) bool {
		if pageSlice[i].Visits == pageSlice[j].Visits {
			return pageSlice[i].URL < pageSlice[j].URL // Sort alphabetically if same count
		}
		return pageSlice[i].Visits > pageSlice[j].Visits
	})
	return pageSlice
}
