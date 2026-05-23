package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	urls := []string{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")

		if !exists {
			return
		}

		parsedHref, err := url.Parse(href)
		if err != nil {
			return
		}

		absoluteURL := baseURL.ResolveReference(parsedHref)

		urls = append(urls, absoluteURL.String())
	})

	return urls, nil

}