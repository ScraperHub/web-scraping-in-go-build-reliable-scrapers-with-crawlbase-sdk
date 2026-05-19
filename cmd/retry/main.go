package main

import (
	"fmt"
	"log"

	"github.com/crawlbase/blog-go-scraping/internal/config"
	"github.com/crawlbase/blog-go-scraping/internal/crawl"
	"github.com/crawlbase/crawlbase-go"
)

func main() {
	api, err := config.NewJSClient()
	if err != nil {
		log.Fatal(err)
	}

	url := config.AmazonProductURL
	opts := map[string]string{
		"scraper":   "amazon-product-details",
		"country":   "US",
		"ajax_wait": "true",
	}

	fmt.Println("Retry with backoff (requires StatusCode 200 and PCStatus 200)")
	res, err := crawl.GetWithBackoff(func() (*crawlbase.Response, error) {
		return api.Get(url, opts)
	}, 3)
	if err != nil {
		if res != nil {
			crawl.PrintLayers("final attempt", res)
		}
		log.Fatal(err)
	}

	crawl.PrintLayers("success", res)
	if res.JSON != nil {
		fmt.Printf("Name: %s\n", crawl.ProductField(res.JSON, "name"))
	}
}
