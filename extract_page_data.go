package main

import (
	"fmt"
	"net/url"
)

func extractPageData(html, pageURL string) PageData {
	pageData := PageData{}

	parsedURL, err := url.Parse(pageURL)
	if err != nil {
		fmt.Println("error:", err)
		return pageData
	}

	headings := getHeadingFromHTML(html)
	firstPara := getFirstParagraphFromHTML(html)
	outgoingLinks, err := getURLsFromHTML(html, parsedURL)
	imageURLs, err := getImagesFromHTML(html, parsedURL)

	pageData.URL = parsedURL.String()
	pageData.Heading = headings
	pageData.FirstParagraph = firstPara
	pageData.OutgoingLinks = outgoingLinks
	pageData.ImageURLs = imageURLs

	return pageData
}