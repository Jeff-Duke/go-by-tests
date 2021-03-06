package concurrency

// WebsiteChecker is a callback fn used to validate a website
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// CheckWebsites checks a list of websites against the provided WebsiteChecker and returns a list of which are valid or invalid
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(url string) {
			resultChannel <- result{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
