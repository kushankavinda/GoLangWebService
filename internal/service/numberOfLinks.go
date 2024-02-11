package service

import (
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func countLinks(url string) (internalCount, externalCount int, err error) {
	// Make an HTTP GET request to the specified URL
	response, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer response.Body.Close()

	// Parse the HTML content of the page
	doc, err := html.Parse(response.Body)
	if err != nil {
		return 0, 0, err
	}

	baseURL := response.Request.URL

	// Count internal and external links
	countLinksRecursive(doc, baseURL, &internalCount, &externalCount)

	return internalCount, externalCount, nil
}

func countLinksRecursive(n *html.Node, baseURL *url.URL, internalCount, externalCount *int) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				linkURL, err := url.Parse(attr.Val)
				if err == nil {
					absoluteURL := baseURL.ResolveReference(linkURL).String()
					if strings.HasPrefix(absoluteURL, baseURL.String()) {
						// Internal link
						(*internalCount)++
					} else {
						// External link
						(*externalCount)++
					}
				}
			}
		}
	}

	// Recursively count links in the child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countLinksRecursive(c, baseURL, internalCount, externalCount)
	}
}
