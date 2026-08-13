package main

import (
	"fmt"
	"log"
	"os"

	"github.com/crawlbase/blog-go-scraping/internal/config"
	"github.com/crawlbase/blog-go-scraping/internal/crawl"
)

func main() {
	normal, js, err := config.MustClients()
	if err != nil {
		log.Fatal(err)
	}

	opts := map[string]string{
		"scraper": "amazon-product-details",
		"country": "US",
	}

	fmt.Println("Attempt 1: normal token + built-in scraper (cheapest path)")
	res, err := normal.Get(config.AmazonProductURL, opts)
	if err != nil {
		log.Fatal(err)
	}
	crawl.PrintLayers("attempt 1", res)

	if crawl.IsCrawlSuccess(res) && res.JSON != nil && crawl.ProductField(res.JSON, "name") != "" {
		printProduct(res.JSON)
		return
	}

	if res.CBStatus == 520 || res.CBStatus == 525 {
		fmt.Println("\nPromoting to JavaScript token with ajax_wait...")
		jsOpts := map[string]string{
			"scraper":   "amazon-product-details",
			"country":   "US",
			"ajax_wait": "true",
		}
		res, err = js.Get(config.AmazonProductURL, jsOpts)
		if err != nil {
			log.Fatal(err)
		}
		crawl.PrintLayers("attempt 2 (JS)", res)
	}

	if !crawl.IsCrawlSuccess(res) {
		fmt.Println("Amazon scrape did not succeed.")
		os.Exit(1)
	}
	if res.JSON == nil {
		log.Fatal("expected JSON from scraper")
	}
	printProduct(res.JSON)
}

func printProduct(data map[string]any) {
	fmt.Printf("\nProduct: %s\n", crawl.ProductField(data, "name"))
	fmt.Printf("Price:   %s\n", crawl.ProductField(data, "price"))
}
