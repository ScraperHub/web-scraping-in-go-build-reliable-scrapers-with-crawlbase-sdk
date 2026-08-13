package crawl

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/crawlbase/crawlbase-go"
)

type Getter func() (*crawlbase.Response, error)

func GetWithBackoff(get Getter, attempts int) (*crawlbase.Response, error) {
	var last *crawlbase.Response
	for i := 0; i < attempts; i++ {
		res, err := get()
		if err != nil {
			return nil, err
		}
		last = res

		if IsCrawlSuccess(res) {
			return res, nil
		}
		if res.StatusCode >= 400 && res.StatusCode < 500 {
			return res, fmt.Errorf("crawlbase client error %d", res.StatusCode)
		}
		if i == attempts-1 {
			break
		}
		d := time.Duration(rand.Float64() * math.Pow(2, float64(i)) * float64(time.Second))
		time.Sleep(d)
	}
	if last == nil {
		return nil, fmt.Errorf("crawl failed after %d attempts", attempts)
	}
	return last, fmt.Errorf("crawl failed: status=%d cb_status=%d", last.StatusCode, last.CBStatus)
}
