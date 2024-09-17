package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("network error: %v", err)
	}

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("client error: %v", res.StatusCode)
	}

	if contentType := res.Header.Get("content-type"); !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("got incorrect content-type, required: 'text/html', got: %v", contentType)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	rawHTML := string(bytes)

	return rawHTML, nil
}
