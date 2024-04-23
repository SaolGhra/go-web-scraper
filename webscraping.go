package main

// Importing necessary packages
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GameInfo struct to represent the game information
type GameInfo struct {
	WebsiteName string   `json:"website_name"`
	Information []string `json:"information"`
}

// saveToJSON saves the game information to a JSON file
func saveToJSON(gameName string, info []GameInfo) {
	data, err := json.MarshalIndent(info, "", " ")
	if err != nil {
		log.Fatalf("Failed to marshal game information: %v\n", err)
		return
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s.json", gameName), data, 0644)

	if err != nil {
		log.Fatalf("Failed to write game information to file: %v\n", err)
		return
	}
}

// Website struct to represent the name, url and selector of the website
type Website struct {
	Name     string
	URL      string
	Selector string
}

func main() {
	// User input for the game name
	var gameName string
	fmt.Print("Enter the game name: ")
	fmt.Scanln(&gameName)

	// Predefined websites with their respective div selectors
	websites := []Website{
		{Name: "STEAM", URL: "https://store.steampowered.com", Selector: ".dev_row:contains('Developer') a"},
		{Name: "IGN", URL: "https://www.ign.com", Selector: ".object-summary-text.developers-info .jsx-153568585.small a"},
		{Name: "GOG", URL: "https://www.gog.com", Selector: ".table__row.details__rating.details__row .details__content table__row-content a"},
		{Name: "RAWG", URL: "https://www.rawg.io", Selector: ".game__meta-block .game__meta-text .game__meta-filter-link itemprop.creator"},
	}

	var info []GameInfo

	// Iterative loop over each website to attempt scraping
	for _, website := range websites {
		information := scrapeWebsite(gameName, Website{URL: website.URL, Selector: website.Selector})
		if information != nil {
			fmt.Printf("Information found on %s: %s\n", website.Name, strings.Join(information, ", "))
			info = append(info, GameInfo{WebsiteName: website.Name, Information: information})
		}
	}

	if len(info) > 0 {
		saveToJSON(gameName, info)
	} else {
		fmt.Println("No information found on any website")
	}

	fmt.Println("No information found on any website")
}

// scrapeWebsite scrapes the website for the given game name using the provided website struct
func scrapeWebsite(gameName string, website Website) []string {
	searchURL := fmt.Sprintf("https://www.google.com/search?q=%s", url.QueryEscape(fmt.Sprintf("%s %s game", gameName, website.Name)))

	searchResponse, err := http.Get(searchURL)
	if err != nil {
		log.Printf("Failed to perform Google search for %s on %s: %v\n", gameName, website.Name, err)
		return nil
	}
	defer searchResponse.Body.Close()

	if searchResponse.StatusCode != http.StatusOK {
		log.Printf("Failed to perform Google search for %s on %s, Status code: %d\n", gameName, website.Name, searchResponse.StatusCode)
		return nil
	}

	searchPage, err := goquery.NewDocumentFromReader(searchResponse.Body)
	if err != nil {
		log.Printf("Failed to parse Google search results for %s on %s: %v\n", gameName, website.Name, err)
		return nil
	}

	websiteLink, exists := searchPage.Find("a[href*='" + website.URL + "']").Attr("href")
	if !exists {
		log.Printf("No %s link found in Google search results\n", website.Name)
		return nil
	}

	websiteURL := extractURL(websiteLink)
	if websiteURL == "" {
		log.Printf("Failed to extract %s URL from Google search results\n", website.Name)
		return nil
	}

	websiteResponse, err := http.Get(websiteURL)
	if err != nil {
		log.Printf("Failed to fetch %s website: %v\n", website.Name, err)
		return nil
	}
	defer websiteResponse.Body.Close()

	if websiteResponse.StatusCode != http.StatusOK {
		log.Printf("Failed to fetch %s website, Status code: %d\n", website.Name, websiteResponse.StatusCode)
		return nil
	}

	websitePage, err := goquery.NewDocumentFromReader(websiteResponse.Body)
	if err != nil {
		log.Printf("Failed to parse %s website: %v\n", website.Name, err)
		return nil
	}

	information := make([]string, 0)
	websitePage.Find(website.Selector).Each(func(i int, s *goquery.Selection) {
		information = append(information, s.Text())
	})

	if len(information) == 0 {
		log.Printf("Information not found on %s\n", website.Name)
		return nil
	}

	return information
}

// extractURL extracts the URL from the Google search result link
func extractURL(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		return ""
	}

	return u.Query().Get("q")
}
