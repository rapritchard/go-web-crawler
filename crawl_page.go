package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error - could not parse base URL '%s': %v\n", rawBaseURL, err)
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - could not parse current URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalisedCurrentURL, err := normaliseURL(currentURL.String())

	if err != nil {
		fmt.Printf("Error - normalisedURL: %v", err)
		return
	}

	if _, ok := pages[normalisedCurrentURL]; ok {
		pages[normalisedCurrentURL]++
		return
	}

	pages[normalisedCurrentURL] = 1

	fmt.Printf("crawling -> %s\n", rawCurrentURL)
	currentHtml, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	links, err := getUrlsFromHTML(currentHtml, rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, link := range links {
		crawlPage(rawBaseURL, link, pages)
	}
}
