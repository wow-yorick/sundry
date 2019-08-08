package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

var (
	m = make(map[string]int)

	l sync.Mutex
	//   群组等待     当添加的任务没有完成时（done（））， wait（） 会一直等待    三个方法   Add()  Done()  Wait()
	i sync.WaitGroup
)

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

func Craw1(url string, depth int, fetcher Fetcher) {
	defer i.Done()
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
	}
	l.Lock()
	if m[url] == 0 {
		m[url]++
		depth--
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			i.Add(1)
			go Craw1(u, depth, fetcher)
		}
	}
	l.Unlock()
	return
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

func main() {
	i.Add(1)
	Craw1("http://golang.org/", 4, fetcher)
	i.Wait()

	for k, _ := range m {
		fmt.Println(k)
	}
	fmt.Println("over")
}
