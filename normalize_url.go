package main

import (
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	u, err := url.Parse(inputURL)
	if err != nil {
		return "",err
	}

	fullURL := u.Host + u.Path 

	fullURL = strings.TrimSuffix(fullURL,"/")

	return fullURL, nil
}