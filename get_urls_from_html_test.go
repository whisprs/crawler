package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute URL",
			inputURL: "https://test.url.dev",
			inputBody: `
<html>
	<body>
		<a href="https://test.url.dev">
			<span>Test URL</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://test.url.dev"},
		},
		{
			name:     "relative URL",
			inputURL: "https://test.url.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Test URL</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://test.url.dev/path/one"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://test.url.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Test URL</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Other Test URL</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://test.url.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no href",
			inputURL: "https://test.url.dev",
			inputBody: `
<html>
	<body>
		<a>
			<span>Test URL</span>
		</a>
	</body>
</html>
`,
			expected: []string{},
		},
		{
			name:     "bad HTML",
			inputURL: "https://test.url.dev",
			inputBody: `
<html body>
	<a href="path/one">
			<span>Test URL</span>
	</a>
</html body>
`,
			expected: []string{"https://test.url.dev/path/one"},
		},
		{
			name:     "invalid href URL",
			inputURL: "https://test.url.dev",
			inputBody: `
<html>
	<body>
		<a href=":\\invalidURL">
			<span>Invalid Test URL</span>
		</a>
	</body>
</html>
`,
			expected: []string{},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - %s FAIL:\n unexpected error: %v", i, tc.name, err)
			}

			result, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Errorf("Test %v - %s FAIL:\n unexpected error: %v", i, tc.name, err)
				return
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Test %v - %s FAIL:\n EXPECTED: %v\n GOT: %v", i, tc.name, tc.expected, result)
			}
		})
	}
}
