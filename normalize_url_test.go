package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "base",
			inputURL: "https://testurl.com",
			expected: "testurl.com",
		},
		{
			name:     "base with path",
			inputURL: "https://testurl.com/path",
			expected: "testurl.com/path",
		},
		{
			name:     "base with path and trailing slash",
			inputURL: "https://testurl.com/path/",
			expected: "testurl.com/path",
		},
		{
			name:     "base with www",
			inputURL: "https://www.testurl.com",
			expected: "testurl.com",
		},
		{
			name:     "base with subdomain",
			inputURL: "https://subdomain.testurl.com",
			expected: "subdomain.testurl.com",
		},
		{
			name:     "base with http",
			inputURL: "http://testurl.com",
			expected: "testurl.com",
		},
		{
			name:     "base with query",
			inputURL: "https://testurl.com/path?query=123",
			expected: "testurl.com/path?query=123",
		},
		{
			name:     "base with path and query",
			inputURL: "https://testurl.com/path?query=123",
			expected: "testurl.com/path?query=123",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - %s FAIL:\n unexpected error: %v", i, tc.name, err)
				return
			}

			if result != tc.expected {
				t.Errorf("Test %v - %s FAIL:\n EXPECTED: %v\n GOT: %v", i, tc.name, tc.expected, result)
			}
		})
	}
}
