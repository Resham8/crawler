package main

import (
	"net/url"
	"reflect"
	"testing"
)

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