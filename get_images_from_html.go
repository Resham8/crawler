package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	
	imageUrls := []string{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, exists := s.Attr("src")

		if !exists {
			return
		}

		parsedSrc, err := url.Parse(src)
		if err != nil {
			return
		}

		absoluteURL := baseURL.ResolveReference(parsedSrc)

		imageUrls = append(imageUrls, absoluteURL.String())
	})

	return imageUrls, nil
}