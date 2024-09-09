package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getUrlsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, fmt.Errorf("could not parse html: %w", err)
	}

	var links []string
	var traverseNodes func(*html.Node)

	traverseNodes = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href, err := url.Parse(a.Val)
					if err != nil {
						fmt.Printf("could not parse href '%v': %v\n", a.Val, err)
						continue
					}
					resolvedURL := baseURL.ResolveReference(href)
					links = append(links, resolvedURL.String())
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverseNodes(c)
		}
	}
	traverseNodes(doc)

	return links, nil
}
