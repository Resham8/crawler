package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}

	h1 := doc.Find("h1").First().Text()
	if h1 != "" {
		return h1
	}

	return doc.Find("h2").First().Text()
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}

	main := doc.Find("main")
	if main.Length() > 0 {
		p := main.Find("p").First()
		if p.Length() > 0 {
			return p.Text()
		}
	}

	pTag := doc.Find("p").First()

	if pTag.Length() > 0 {
		return pTag.Text()
	}

	return ""
}


func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	urls := []string {}

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

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	imageUrls := []string {}

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