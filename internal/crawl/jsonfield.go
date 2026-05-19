package crawl

import "fmt"

func ScraperBody(data map[string]any) map[string]any {
	if data == nil {
		return nil
	}
	if body, ok := data["body"].(map[string]any); ok {
		return body
	}
	return data
}

func JSONString(data map[string]any, key string) string {
	v, ok := data[key]
	if !ok || v == nil {
		return ""
	}
	switch s := v.(type) {
	case string:
		return s
	default:
		return fmt.Sprint(v)
	}
}

func ProductField(data map[string]any, key string) string {
	return JSONString(ScraperBody(data), key)
}
