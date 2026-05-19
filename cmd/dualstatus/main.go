package main

import (
	"fmt"
	"log"
	"os"

	"github.com/crawlbase/blog-go-scraping/internal/config"
	"github.com/crawlbase/blog-go-scraping/internal/crawl"
	"github.com/crawlbase/crawlbase-go"
)

const (
	notFoundURL      = "https://httpbin.org/status/404"
	invalidAmazonURL = "https://www.amazon.com/dp/0000000000"
)

func main() {
	api, err := config.NewNormalClient()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Case 1: target returns HTTP 404 (httpbin) — Crawlbase still answers 200")
	res404, err := api.Get(notFoundURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	crawl.PrintLayers("httpbin 404", res404)
	warnLayers(res404)

	fmt.Println("\nCase 2: invalid Amazon ASIN — expect PCStatus 404")
	resBad, err := api.Get(invalidAmazonURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	crawl.PrintLayers("invalid Amazon ASIN", resBad)
	warnLayers(resBad)

	fmt.Println("\nCase 3: Amazon with normal token only (no scraper) — may get 520/525")
	resAmazon, err := api.Get(config.AmazonProductURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	crawl.PrintLayers("amazon normal token", resAmazon)
	warnLayers(resAmazon)

	demonstrated := false
	for _, res := range []*crawlbase.Response{res404, resBad, resAmazon} {
		if crawl.NaiveSuccess(res) && (!crawl.IsCrawlSuccess(res) || res.OriginalStatus >= 400) {
			demonstrated = true
		}
	}
	if demonstrated {
		fmt.Println("\nDemonstrated: StatusCode 200 does not mean a usable page — check PCStatus and OriginalStatus.")
		os.Exit(0)
	}
	fmt.Println("\nAll three returned clean PCStatus — layers were still printed for inspection.")
}

func warnLayers(res *crawlbase.Response) {
	if !crawl.NaiveSuccess(res) {
		return
	}
	if !crawl.IsCrawlSuccess(res) {
		fmt.Println("WARNING: StatusCode 200 but PCStatus != 200 — naive check would hide the failure.")
	}
	if res.OriginalStatus >= 400 {
		fmt.Println("WARNING: StatusCode and PCStatus are 200, but OriginalStatus shows a site-side error.")
	}
}
