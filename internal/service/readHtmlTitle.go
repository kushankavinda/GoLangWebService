package service

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func getPageTitle(url string) (string, error) {
	// Make an HTTP GET request to the specified URL
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Parse the HTML content of the page
	doc, err := html.Parse(response.Body)
	if err != nil {
		return "", err
	}

	// Find the title element and extract its text content
	title, found := findTitle(doc)
	if !found {
		return "", fmt.Errorf("Title not found on the page")
	}

	return title, nil
}

func findTitle(n *html.Node) (string, bool) {
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if title, found := findTitle(c); found {
			return title, true
		}
	}

	return "", false
}
