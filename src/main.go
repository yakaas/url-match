package main

import (
	"fmt"
	"net/url"
)

func main() {

	var urls []string = []string{
		"http://whereis.seadonkey.com/api/water",
		"https://somewhere.in.the.sea/api/rocks",
		"https://somewhere.in.the.sea:9898/api/rocks",
	}

	for _, url := range urls {
		parsed := parseUrl(url)
		fmt.Println(url, parsed)
	}

}

func parseUrl(sourceUrl string) string {

	if len(sourceUrl) < 1 {
		return ""
	}

	u, err := url.Parse(sourceUrl)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Source: %s\n", sourceUrl)
	fmt.Printf("Scheme: %s\n", u.Scheme)
	fmt.Printf("Host: %s\n", u.Host)
	fmt.Printf("Path: %s\n", u.Path)
	fmt.Println()

	return sourceUrl
}
