package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetHeadingFromHTMLBasic(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name:      "returns h1 title",
			inputBody: "<html><body><h1>Test Title</h1></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "prefers h1 over h2",
			inputBody: "<html><body><h1>Test Title</h1><h2>Test h2 Title</h2></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "falls back to h2",
			inputBody: "<html><body><h2>Test h2 Title</h2></body></html>",
			expected:  "Test h2 Title",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFromHTML(tc.inputBody)

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name: "first paragraph",
			inputBody: `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name: "paragraph",
			inputBody: `<html><body>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name: "just paragraph",
			inputBody: `<html><body>
		<p>Outside paragraph.</p>
	</body></html>`,
			expected: "Outside paragraph.",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetURLsFromHTMLAbsolute(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "absolute URL",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="https://other.com/page">Link</a></body></html>`,
			expected:  []string{"https://other.com/page"},
		},
		{
			name:      "relative URL",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="/about">About</a></body></html>`,
			expected:  []string{"https://crawler-test.com/about"},
		},
		{
			name:      "multiple links",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="/a">A</a><a href="/b">B</a></body></html>`,
			expected:  []string{"https://crawler-test.com/a", "https://crawler-test.com/b"},
		},
		{
			name:      "no links",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><p>nothing here</p></body></html>`,
			expected:  []string{},
		},
		{
			name:      "malformed href is skipped",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="://bad">bad</a><a href="/good">good</a></body></html>`,
			expected:  []string{"https://crawler-test.com/good"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("couldn't parse input URL: %v", err)
				return
			}

			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

}

func TestGetImagesFromHTMLRelative(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "relative src",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://crawler-test.com/logo.png"},
		},
		{
			name:      "absolute src",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="https://cdn.example.com/photo.jpg"></body></html>`,
			expected:  []string{"https://cdn.example.com/photo.jpg"},
		},
		{
			name:      "multiple images",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="/a.png"><img src="/b.png"></body></html>`,
			expected:  []string{"https://crawler-test.com/a.png", "https://crawler-test.com/b.png"},
		},
		{
			name:      "no images",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><p>no images here</p></body></html>`,
			expected:  []string{},
		},
		{
			name:      "malformed src is skipped",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="://bad"><img src="/good.png"></body></html>`,
			expected:  []string{"https://crawler-test.com/good.png"},
		},
		{
			name:      "img without src is skipped",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img alt="no src"><img src="/valid.png"></body></html>`,
			expected:  []string{"https://crawler-test.com/valid.png"},
		},
		{
			name:      "nested image inside anchor",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="/page"><img src="/thumb.png"></a></body></html>`,
			expected:  []string{"https://crawler-test.com/thumb.png"},
		},
		{
			name:      "src with path traversal resolves correctly",
			inputURL:  "https://crawler-test.com/blog/post/",
			inputBody: `<html><body><img src="../../assets/hero.png"></body></html>`,
			expected:  []string{"https://crawler-test.com/assets/hero.png"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("couldn't parse input URL: %v", err)
				return
			}

			actual, err := getImagesFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
