package main

import (
	"fmt"
	"log"
	"os"

	"github.com/crawlbase/blog-go-scraping/internal/config"
	"github.com/crawlbase/blog-go-scraping/internal/crawl"
)

func main() {
	api, err := config.NewJSClient()
	if err != nil {
		log.Fatal(err)
	}

	res, err := api.Get(config.AmazonProductURL, map[string]string{
		"scraper":   "amazon-product-details",
		"country":   "US",
		"ajax_wait": "true",
	})
	if err != nil {
		log.Fatal(err)
	}

	crawl.PrintLayers("amazon-product-details scraper", res)
	if !crawl.IsCrawlSuccess(res) {
		fmt.Println("Scraper call failed — check PCStatus; retry with JS token if 520/525.")
		os.Exit(1)
	}

	if res.JSON == nil {
		log.Fatal("expected JSON body from built-in scraper")
	}

	fmt.Printf("Name:     %s\n", crawl.ProductField(res.JSON, "name"))
	fmt.Printf("Price:    %s\n", crawl.ProductField(res.JSON, "price"))
	fmt.Printf("Currency: %s\n", crawl.ProductField(res.JSON, "currency"))
	fmt.Printf("Rating:   %s\n", crawl.ProductField(res.JSON, "customerReview"))
}
