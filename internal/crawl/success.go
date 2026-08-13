package crawl

import "github.com/crawlbase/crawlbase-go"

func IsCrawlSuccess(res *crawlbase.Response) bool {
	return res.StatusCode == 200 && res.CBStatus == 200
}

func NaiveSuccess(res *crawlbase.Response) bool {
	return res.StatusCode == 200
}
