package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %v", err)
	}

	hostname := strings.TrimPrefix(parsedURL.Hostname(), "www.")

	fullPath := hostname + parsedURL.Path

	if parsedURL.RawQuery != "" {
		fullPath += "?" + parsedURL.RawQuery
	}

	fullPath = strings.ToLower(fullPath)

	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}
