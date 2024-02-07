package service

import (
	"net/http"

	"golang.org/x/net/html"
)

func countHeadings(url string) (int, error) {
	// Make an HTTP GET request to the specified URL
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	// Parse the HTML content of the page
	doc, err := html.Parse(response.Body)
	if err != nil {
		return 0, err
	}

	// Count the number of headings (h1 to h6)
	headingsCount := 0
	countHeadingsRecursive(doc, &headingsCount)

	return headingsCount, nil
}

func countHeadingsRecursive(n *html.Node, count *int) {
	if n.Type == html.ElementNode && (n.Data == "h1" || n.Data == "h2" || n.Data == "h3" || n.Data == "h4" || n.Data == "h5" || n.Data == "h6") {
		(*count)++
	}

	// Recursively count headings in the child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countHeadingsRecursive(c, count)
	}
}
