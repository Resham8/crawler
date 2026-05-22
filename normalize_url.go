package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	u, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	fullURL := u.Host + u.Path 

	fullURL = strings.TrimSuffix(fullURL,"/")

	return fullURL, nil
}