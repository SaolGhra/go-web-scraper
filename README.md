# Web Scraping in Go

![GitHub Repo stars](https://img.shields.io/github/stars/SaolGhra/go-web-scraper)

This is a simple web scraping program written in Go. It takes a game name as input, scrapes information about the game from various websites, and saves the information to a JSON file.

## How it Works

The program uses predefined websites and their respective CSS selectors to scrape information. The websites and selectors are defined in the `websites` slice in the `main` function:

```go
websites := []Website{
	{Name: "STEAM", URL: "https://store.steampowered.com", Selector: ".dev_row:contains('Developer') a"},
	{Name: "IGN", URL: "https://www.ign.com", Selector: ".object-summary-text.developers-info .jsx-153568585.small a"},
	{Name: "GOG", URL: "https://www.gog.com", Selector: ".table__row.details__rating.details__row .details__content table__row-content a"},
	{Name: "RAWG", URL: "https://www.rawg.io", Selector: ".game__meta-block .game__meta-text .game__meta-filter-link itemprop.creator"},
}
```

Each `Website` struct contains the name of the website, the URL to scrape, and the CSS selector to find the information.

The program is powerful and fast. It can search multiple websites simultaneously and save the information quickly.

## Customization

You can customize the program to scrape information from other websites. All you need to do is change the `Name`, `URL`, and `Selector` fields in the `Website` structs.

For example, if you want to scrape information from a website called "Example", you would add a `Website` struct like this:

```go
{Name: "Example", URL: "https://www.example.com", Selector: ".example-selector"}
```

Please note that you need to find the correct CSS selector for the information you want to scrape. You can do this by inspecting the webpage in your browser.

## Imports

The program uses the following packages:

- `fmt`
- `net/http`
- `github.com/PuerkitoBio/goquery`
- `strings`
- `encoding/json`
- `io/ioutil`
- `log`

Please make sure to install the `goquery` package before running the program. You can do this with the following command:

```bash
go get github.com/PuerkitoBio/goquery
```

## Disclaimer

Please note that web scraping should be done in accordance with the website's terms of service. Some websites explicitly disallow web scraping in their robots.txt file or Terms of Service. Always respect the website's rules and the privacy of its users.
