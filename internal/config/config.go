package config

import (
	"fmt"
	"os"

	"github.com/crawlbase/crawlbase-go"
)

const AmazonProductURL = "https://www.amazon.com/Apple-Version-Unlocked-Renewed-Premium/dp/B0G45YW56F/ref=sr_1_2"

func NormalToken() (string, error) {
	t := os.Getenv("CRAWLBASE_TOKEN")
	if t == "" {
		return "", fmt.Errorf("CRAWLBASE_TOKEN is not set")
	}
	return t, nil
}

func JSToken() (string, error) {
	t := os.Getenv("CRAWLBASE_JS_TOKEN")
	if t == "" {
		return "", fmt.Errorf("CRAWLBASE_JS_TOKEN is not set")
	}
	return t, nil
}

func NewNormalClient() (*crawlbase.CrawlingAPI, error) {
	token, err := NormalToken()
	if err != nil {
		return nil, err
	}
	return crawlbase.NewCrawlingAPI(token)
}

func NewJSClient() (*crawlbase.CrawlingAPI, error) {
	token, err := JSToken()
	if err != nil {
		return nil, err
	}
	return crawlbase.NewCrawlingAPI(token)
}

func MustClients() (*crawlbase.CrawlingAPI, *crawlbase.CrawlingAPI, error) {
	normal, err := NewNormalClient()
	if err != nil {
		return nil, nil, err
	}
	js, err := NewJSClient()
	if err != nil {
		return nil, nil, err
	}
	return normal, js, nil
}
