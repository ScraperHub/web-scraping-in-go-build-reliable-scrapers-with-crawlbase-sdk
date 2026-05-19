package main

import (
	"fmt"
	"log"

	"github.com/crawlbase/blog-go-scraping/internal/config"
	"github.com/crawlbase/blog-go-scraping/internal/crawl"
)

func main() {
	normal, js, err := config.MustClients()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Normal token client — static HTML, built-in scrapers, faster/cheaper.")
	res, err := normal.Get("https://example.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	crawl.PrintLayers("normal token / example.com", res)

	fmt.Println("\nJavaScript token client — browser rendering (page_wait, ajax_wait, scroll).")
	resJS, err := js.Get(config.AmazonProductURL, map[string]string{
		"ajax_wait": "true",
	})
	if err != nil {
		log.Fatal(err)
	}
	crawl.PrintLayers("JS token / Amazon", resJS)

	fmt.Println("\nKeep two clients when you alternate tokens; the SDK does not switch per call.")
}
