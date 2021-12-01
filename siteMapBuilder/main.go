package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"siteMapBuilder/parseWebPage"
	"strings"
)

func main() {
	// set up flags
	urlFlag := flag.String("url", "https://gophercises.com", "url that you want to build sitemap for")
	depth := flag.Int("depth", 10, "the maximum depth for number of links to traverse")
	flag.Parse()

	pages := bfs(*urlFlag, *depth)
	//pages := get(*urlFlag)
	for _, l := range pages {
		fmt.Println(l)
	}

}

// bfs does a breadth first search on the given url and returns the links founds
// maxDepth denotes how deep the search will go
func bfs(urlString string, maxDepth int) []string {
	// seen stores list of urls seen
	seen := make(map[string]struct{})
	var q map[string]struct{}
	// var nq stores the url to be searched at first
	nq := map[string]struct{}{
		urlString: {},
	}

	// starting the bfs for given depth
	for i := 0; i <= maxDepth; i++ {
		// adding nextqueue(nq) to q
		q, nq = nq, make(map[string]struct{})
		if len(q) == 0 {
			break
		}
		// range over q to get more links
		for url := range q {
			// if url is already visited, skip
			if _, ok := seen[url]; ok {
				continue
			}
			// add new url to seen list
			seen[url] = struct{}{}
			// add links inside new url to next queue
			for _, link := range get(url) {
				nq[link] = struct{}{}
			}
		}
	}

	// make []string out of seen links
	ret := make([]string, 0, len(seen))
	for url := range seen {
		ret = append(ret, url)
	}
	return ret
}

// get takes in a url and returns all the links from the same domain
func get(urlStr string) []string {
	// get web page
	resp, err := http.Get(urlStr)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()

	//create base url
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	return filter(hrefs(resp.Body, base), withPrefix(base))
}

// hrefs takes in the webpage and the base url
// returns a list of urls found on the webpage
func hrefs(body io.Reader, base string) []string {
	// parse links on the webpage
	links, _ := parseWebPage.Parse(body)
	//filter and fix parsed links
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
