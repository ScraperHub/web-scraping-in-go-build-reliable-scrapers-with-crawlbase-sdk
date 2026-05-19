package main

import (
	"fmt"
	"log"
	"os"

	"github.com/crawlbase/blog-go-scraping/internal/config"
	"github.com/crawlbase/blog-go-scraping/internal/crawl"
)

func main() {
	api, err := config.NewNormalClient()
	if err != nil {
		log.Fatal(err)
	}

	res, err := api.Get("https://example.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	crawl.PrintLayers("example.com", res)
	if crawl.IsCrawlSuccess(res) {
		fmt.Printf("Body length: %d bytes\n", len(res.Body))
		os.Exit(0)
	}
	fmt.Println("Crawl did not succeed — check PCStatus above.")
	os.Exit(1)
}
