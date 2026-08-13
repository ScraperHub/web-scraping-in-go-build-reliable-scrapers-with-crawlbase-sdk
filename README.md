<a href="https://crawlbase.com/signup?utm_source=github&utm_medium=readme&utm_campaign=crawling_api_banner" target="_blank">
  <img src="https://github.com/user-attachments/assets/afa4f6e7-25fb-442c-af2f-b4ddcfd62ab2" 
       alt="crawling-api-cta" 
       style="max-width: 100%; border: 0;">
</a>



# Web Scraping in Go with Crawlbase

Runnable examples for the blog post **Web Scraping in Go: Build Reliable Scrapers with Crawlbase SDK**. Each command under `cmd/` maps to one section of the article.

## Prerequisites

- Go 1.22 or newer
- A [Crawlbase](https://crawlbase.com/signup) account with **Normal** and **JavaScript** tokens (1,000 free requests, no credit card)

## Install

```powershell
cd code
go mod download
```

The module already depends on [github.com/crawlbase/crawlbase-go](https://github.com/crawlbase/crawlbase-go). To add it in your own project:

```powershell
go get github.com/crawlbase/crawlbase-go
```

## Set your tokens

Do not commit API tokens. Export them in your shell:

PowerShell:

```powershell
$env:CRAWLBASE_TOKEN = "YOUR_NORMAL_TOKEN"
$env:CRAWLBASE_JS_TOKEN = "YOUR_JAVASCRIPT_TOKEN"
```

macOS or Linux:

```bash
export CRAWLBASE_TOKEN="YOUR_NORMAL_TOKEN"
export CRAWLBASE_JS_TOKEN="YOUR_JAVASCRIPT_TOKEN"
```

## Run the examples

| Command | What it demonstrates |
|---------|----------------------|
| `go run ./cmd/quickstart` | First crawl with the Normal token |
| `go run ./cmd/dualstatus` | Why `StatusCode == 200` is not enough |
| `go run ./cmd/tokens` | Normal vs JavaScript clients |
| `go run ./cmd/jsrender` | JS rendering with `ajax_wait` on Amazon |
| `go run ./cmd/scraper` | Built-in `amazon-product-details` scraper |
| `go run ./cmd/amazon` | Geo + scraper + promote to JS on 520/525 |
| `go run ./cmd/retry` | Backoff until `StatusCode` and `CBStatus` are 200 |

Amazon product URL used in examples:

`https://www.amazon.com/Apple-Version-Unlocked-Renewed-Premium/dp/B0G45YW56F/ref=sr_1_2`

## Dual status codes

Crawlbase returns three signals on every response:

1. **`StatusCode`** — HTTP status of your request to Crawlbase (`200` means Crawlbase processed the call).
2. **`CBStatus`** — Crawlbase’s verdict on the target page (`200` means a clean fetch).
3. **`OriginalStatus`** — HTTP status the target site returned (e.g. `404` on a missing page).

Always check `StatusCode` first, then `CBStatus`, then `OriginalStatus` for site-side errors. See the [status codes docs](https://crawlbase.com/docs/status-codes) and [Crawling API](https://crawlbase.com/docs/crawling-api).

Helper functions live in `internal/crawl`:

- `IsCrawlSuccess(res)` — `StatusCode == 200 && CBStatus == 200`
- `PrintLayers(label, res)` — prints all three layers for debugging

## Troubleshooting

| Symptom | Likely cause | What to try |
|---------|--------------|-------------|
| `StatusCode 200` but empty or wrong body | Only checked HTTP status | Check `CBStatus` and `OriginalStatus` |
| `CBStatus 404` | Target page missing | Fix URL or ASIN |
| `CBStatus 520` or `525` | Empty body or unsolved bot challenge | JavaScript token, `ajax_wait`, or `page_wait` |
| `StatusCode 403` | JS-only option on Normal token | Use the JavaScript token |
| `StatusCode 429` | Concurrency limit | Back off and retry |
| `StatusCode 401` or `402` | Invalid token or no credits | Verify token and account balance |

Failed crawls (`CBStatus != 200`) do not count against your quota; retries are free until you get a successful `CBStatus`.

## Alternatives

- **[Colly](https://github.com/gocolly/colly)** — excellent for fast, static HTML crawling in Go; you still own proxies, JavaScript rendering, and anti-bot handling.
- **[chromedp](https://github.com/chromedp/chromedp)** — full headless Chrome locally; powerful but heavy to operate at scale.
- **Crawlbase** — managed crawling (proxies, JS rendering, scrapers) via one API and the [Go SDK](https://github.com/crawlbase/crawlbase-go).

## Links

- [Free Crawlbase signup](https://crawlbase.com/signup) — 1,000 requests, no credit card
- [crawlbase-go on GitHub](https://github.com/crawlbase/crawlbase-go)
- [Go SDK docs](https://crawlbase.com/docs/sdk-go)
