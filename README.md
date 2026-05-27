# Crawler

A concurrent web crawler written in Go. Given a starting URL, it recursively discovers and visits all pages within the same domain, extracting structured page data and writing a JSON report of the results.

## Features

- **Concurrent crawling** — configurable number of goroutines crawl pages in parallel
- **Domain-scoped** — only follows links within the same hostname as the seed URL
- **Page limit** — stops after a configurable maximum number of pages
- **Data extraction** — for each page, captures the heading, first paragraph, outgoing links, and image URLs
- **JSON report** — writes results to `report.json` on completion
- **URL normalization** — deduplicates URLs to avoid revisiting the same page

## Requirements

- Go 1.26

## Installation

```bash
git clone https://github.com/Resham8/crawler.git
cd crawler
go build .
```

## Running Tests

```bash
go test ./...
```

## Usage

```bash
./crawler <URL> <maxConcurrency> <maxPages>
```

| Argument         | Description                                            |
| ---------------- | ------------------------------------------------------ |
| `URL`            | The starting URL to crawl (e.g. `https://example.com`) |
| `maxConcurrency` | Maximum number of pages to crawl concurrently          |
| `maxPages`       | Maximum total number of pages to crawl                 |

**Example:**

```bash
./crawler https://example.com 5 100
```

This starts crawling `https://example.com` using up to 5 concurrent goroutines and stops after 100 pages are visited.

## Output

While running, the crawler prints each URL as it is discovered:

```
starting crawl of: https://example.com...
crawling https://example.com
crawling https://example.com/about
crawling https://example.com/blog
...
found: example.com/about
found: example.com/blog
...
```

On completion, a `report.json` file is written containing the collected data for every crawled page:

```json
{
  "example.com/about": {
    "url": "https://example.com/about",
    "heading": "About Us",
    "first_paragraph": "We are a ...",
    "outgoing_links": [
      "https://example.com/team",
      "https://example.com/contact"
    ],
    "image_urls": ["https://example.com/images/hero.png"]
  }
}
```
