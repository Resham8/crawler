package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseUrl, err := url.Parse(rawBaseURL)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	currentUrl, err := url.Parse(rawCurrentURL)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	if baseUrl.Host != currentUrl.Host {
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}

	if pages[normalizedCurrentURL] > 0 {
		pages[normalizedCurrentURL]++
		return
	} else {
		pages[normalizedCurrentURL] = 1
	}

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("got HTML for: %s\n", rawCurrentURL)

	links, err := getURLsFromHTML(htmlBody, baseUrl)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, link := range links {		
		crawlPage(rawBaseURL, link, pages)
		// fmt.Println("Html response: ")
		// fmt.Println(htmlBody)
	}

}
