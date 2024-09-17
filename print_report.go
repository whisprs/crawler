package main

import (
	"fmt"
	"sort"
)

type Page struct {
	url   string
	count int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("REPORT for %v\n", baseURL)
	fmt.Println("=============================")

	sorted := sortPages(pages)

	for _, page := range sorted {
		fmt.Printf("Found %d internal links to %s\n", page.count, page.url)
	}
}

// TODO proper sorting
func sortPages(pages map[string]int) []Page {
	report := make([]Page, 0, len(pages))
	for k, v := range pages {
		report = append(report, Page{k, v})
	}

	sort.SliceStable(report, func(i, j int) bool {
		return report[i].count < report[j].count
	})

	return report
}
