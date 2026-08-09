// Exercise: Web Crawler
// Link: go.dev/tour/concurrency/10

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *SafeCache, wg *sync.WaitGroup) {
	// Issue: Doesn't cache the value. Just keeps track of each URL being processed.
	// Individual caches, i.e., Get & Has methods will have a race condition where
	// goroutines might all clear Has because there is no said `url` in cache yet.
	if depth <= 0 || cache.isCached(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		wg.Go(func() { Crawl(u, depth-1, fetcher, cache, wg) })
	}
}

func main() {
	cache := SafeCache{m: make(map[string]bool)}
	var wg sync.WaitGroup

	wg.Go(func() { Crawl("https://golang.org/", 4, fetcher, &cache, &wg) })

	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

type SafeCache struct {
	mu sync.Mutex
	m  map[string]bool
}

func (c *SafeCache) isCached(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.m[url]; exists {
		return true
	}

	c.m[url] = true
	return false
}

func (c *SafeCache) Has(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, exists := c.m[url]
	return exists
}

func (c *SafeCache) Set(url string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[url] = true
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
