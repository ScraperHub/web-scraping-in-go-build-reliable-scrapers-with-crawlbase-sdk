package crawl

import (
	"fmt"

	"github.com/crawlbase/crawlbase-go"
)

func PrintLayers(label string, res *crawlbase.Response) {
	fmt.Printf("\n--- %s ---\n", label)
	fmt.Printf("StatusCode (Crawlbase API):     %d\n", res.StatusCode)
	fmt.Printf("PCStatus (target verdict):      %d\n", res.PCStatus)
	fmt.Printf("OriginalStatus (site HTTP):     %d\n", res.OriginalStatus)
	if res.URL != "" {
		fmt.Printf("Resolved URL:                   %s\n", res.URL)
	}
}
