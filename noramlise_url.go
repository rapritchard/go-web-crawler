package main

import (
	"fmt"
	"net/url"
	"path"
	"regexp"
	"strings"
)

func normaliseURL(rawURL string) (string, error) {
	re := regexp.MustCompile(`^(https?:)//+`)
	processedURL := re.ReplaceAllString(rawURL, "$1//")

	parsedURL, err := url.Parse(strings.ToLower(processedURL))
	if err != nil {
		return "", fmt.Errorf("could not parse URL: %w", err)
	}

	normalisedURL := parsedURL.Hostname()
	cleanedPath := path.Clean(parsedURL.EscapedPath())
	if cleanedPath != "/" && cleanedPath != "." {
		normalisedURL += cleanedPath
	} else {
		normalisedURL += "/"
	}

	return normalisedURL, nil
}
