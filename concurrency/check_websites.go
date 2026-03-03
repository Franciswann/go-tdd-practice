package concurrency

// WebsiteChecker is a function type that checks if a website is accessiable.
// It takes a URL string and returns true if the website is accessiable, false otherwise.
type WebsiteChecker func(string) bool

// result represents the result of checking a single website.
// The string field holds the URL, and the bool field holds the accessibility status.
type result struct {
	string
	bool
}

// CheckWebsites checks multiple URLs concurrently using goroutines and channels.
// It launches a goroutine for each URL and collect results through a channel.
// Returns a map mapping each URL to its accessibility result.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			// send statement
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		// receive expression
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
