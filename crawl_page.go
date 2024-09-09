package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{} // acquire a slot
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}() // release a slot

	if cfg.pagesLen() >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: could not parse current URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalisedCurrentURL, err := normaliseURL(currentURL.String())
	if err != nil {
		fmt.Printf("Error - normalisedURL: %v", err)
		return
	}

	isFirst := cfg.addPageVisit(normalisedCurrentURL)

	if !isFirst {
		return
	}

	fmt.Printf("crawling -> %s\n", rawCurrentURL)
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	links, err := getUrlsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, link := range links {
		cfg.wg.Add(1)
		go cfg.crawlPage(link)
	}
}
