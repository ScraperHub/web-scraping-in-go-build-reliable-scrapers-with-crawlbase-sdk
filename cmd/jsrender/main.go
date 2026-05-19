package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/crawlbase/blog-go-scraping/internal/config"
	"github.com/crawlbase/blog-go-scraping/internal/crawl"
)

func main() {
	api, err := config.NewJSClient()
	if err != nil {
		log.Fatal(err)
	}

	res, err := api.Get(config.AmazonProductURL, map[string]string{
		"ajax_wait": "true",
		"country":   "US",
	})
	if err != nil {
		log.Fatal(err)
	}

	crawl.PrintLayers("Amazon JS render", res)
	if !crawl.IsCrawlSuccess(res) {
		fmt.Println("JS render did not succeed — try ajax_wait or page_wait tuning.")
		os.Exit(1)
	}

	title := extractTitle(res.Body)
	if title != "" {
		fmt.Printf("Page title snippet: %s\n", title)
	} else {
		fmt.Printf("Body length: %d bytes\n", len(res.Body))
	}
}

func extractTitle(html string) string {
	const open = "<title>"
	const close = "</title>"
	i := strings.Index(strings.ToLower(html), open)
	if i < 0 {
		return ""
	}
	i += len(open)
	j := strings.Index(strings.ToLower(html[i:]), close)
	if j < 0 {
		return ""
	}
	return strings.TrimSpace(html[i : i+j])
}
