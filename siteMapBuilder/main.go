package main

import (
	"flag"
	"fmt"
	"net/http"
	"siteMapBuilder/parseWebPage"
)

func main() {
	// set up flags
	urlFlag := flag.String("url", "https://gophercises.com", "url that you want to build sitemap for")
	flag.Parse()

	// get web page
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// parse links on the webpage
	links, _ := parseWebPage.Parse(resp.Body)
	for _, l := range links {
		fmt.Println(l)
	}
	// build proper url with our links

	// filter out links with different domain name

	// find all the pages with Breadth First Search

	// print out xml

}
