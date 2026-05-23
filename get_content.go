package main

import (
	"fmt"	
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

