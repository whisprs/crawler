package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("not enought arguments")
		os.Exit(1)
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := args[0]

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Error converting maxConcurrency:", err)
		return
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error converting maxPages:", err)
		return
	}

	arguments := struct {
		baseURL        string
		maxConcurrency int
		maxPages       int
	}{
		baseURL:        rawBaseURL,
		maxConcurrency: maxConcurrency,
		maxPages:       maxPages,
	}

	parsedURL, err := url.Parse(arguments.baseURL)
	if err != nil {
		fmt.Printf("invalid base url: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("starting crawl %s\n", rawBaseURL)

	cfg := config{
		pages:              make(map[string]int),
		baseURL:            parsedURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, arguments.maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           arguments.maxPages,
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, cfg.baseURL.String())
}
