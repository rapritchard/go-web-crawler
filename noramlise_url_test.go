package main

import (
	"strings"
	"testing"
)

func TestNormaliseURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
		errorContains string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "case sensitivity",
			inputURL: "http://Blog.Boot.Dev/Path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "empty path",
			inputURL: "https://blog.boot.dev",
			expected: "blog.boot.dev/",
		},
		{
			name:     "empty path with a hanging slash",
			inputURL: "https://blog.boot.dev/",
			expected: "blog.boot.dev/",
		},
		{
			name:     "multiple trailing slashes",
			inputURL: "https:///blog.boot.dev///path//to//page///",
			expected: "blog.boot.dev/path/to/page",
		},
		{
			name:     "no schema",
			inputURL: "www.blog.boot.dev/path/",
			expected: "www.blog.boot.dev/path",
		},
		{
			name:          "handle invalid URL",
			inputURL:      `:\\invalidURL`,
			expected:      "",
			errorContains: "could not parse URL",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normaliseURL(tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
