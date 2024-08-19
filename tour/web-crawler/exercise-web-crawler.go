package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type UrlCache struct {
	mu      sync.Mutex
	fetched map[string]bool
}

func (cache *UrlCache) Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	cache.mu.Lock()
	if cache.fetched[url] {
		cache.mu.Unlock()
		return
	}
	cache.fetched[url] = true
	cache.mu.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		cache.mu.Lock()
		if !cache.fetched[u] {
			wg.Add(1)
			go cache.Crawl(u, depth-1, fetcher, wg)
		}
		cache.mu.Unlock()
	}
}

func main() {
	cache := UrlCache{fetched: make(map[string]bool)}
	var wg sync.WaitGroup

	wg.Add(1)
	go cache.Crawl("https://golang.org/", 4, fetcher, &wg)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

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
