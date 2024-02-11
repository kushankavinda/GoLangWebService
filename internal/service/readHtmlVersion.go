package service

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func getHTMLVersion(url string) (string, error) {
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

	// Find the <!DOCTYPE> declaration and extract the HTML version
	htmlVersion := findHTMLVersion(doc)
	if htmlVersion == "" {
		return "HTML version not specified", nil
	}

	return htmlVersion, nil
}

func findHTMLVersion(n *html.Node) string {
	if n.Type == html.DocumentNode && n.FirstChild != nil && n.FirstChild.Type == html.DoctypeNode {
		doctype := n.FirstChild.Data
		if strings.Contains(doctype, "HTML 4.01") {
			return "HTML 4.01"
		} else if strings.Contains(doctype, "HTML 5") {
			return "HTML 5"
		}
		// Add more conditions for other HTML versions as needed
	}

	// Recursively search for the <!DOCTYPE> declaration
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		version := findHTMLVersion(c)
		if version != "" {
			return version
		}
	}

	return ""
}
