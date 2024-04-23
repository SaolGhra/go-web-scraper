package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func scrapeSite(url string) (string, error) {
	response, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(bytes), nil
}

func findDiv(content, divID string) (string, error) {
	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		return "", err
	}
	var find func(*html.Node) string
	find = func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == divID {
					var sb strings.Builder
					html.Render(&sb, n)
					return sb.String()
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if found := find(c); found != "" {
				return found
			}
		}
		return ""
	}
	return find(doc), nil
}

func main() {
	sites := map[string]string{
		"https://www.saolghra.co.uk": ".projects-title",
		"https://www.yahoo.com":      "main",
		"https://www.bing.com":       "main",
		"https://www.ign.com":        "section",
	}

	for site, divID := range sites {
		contnet, err := scrapeSite(site)

		if err != nil {
			fmt.Printf("Error scraping site %s: %s\n", site, err)
			continue
		}

		divContent, err := findDiv(contnet, divID)

		if err != nil {
			fmt.Printf("Error finding div '%s' on %s: %v\n", divID, site, err)
			continue
		}

		if divContent != "" {
			fmt.Printf("Found div '%s' on %s\n", divID, site)
			break
		} else {
			fmt.Printf("Did not find div '%s' on %s\n", divID, site)
		}
	}
	fmt.Println("Scraping Complete")
}
