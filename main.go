package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(os.Args) < 4 {
		fmt.Println("usage: crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]
	maxConcurrency := os.Args[2]
	val, err := strconv.Atoi(maxConcurrency)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	maxPages := os.Args[3]

	maxPagesVal, err := strconv.Atoi(maxPages)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cfg, err := configure(rawBaseURL, val)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	cfg.maxPages = maxPagesVal

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for normalizedURL := range cfg.pages {
		fmt.Printf("found: %s\n", normalizedURL)
	}
	writeJSONReport(cfg.pages, "report.json")
}
