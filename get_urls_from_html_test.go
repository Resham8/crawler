package main

import (
	"net/url"
	"reflect"
	"testing"
)

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