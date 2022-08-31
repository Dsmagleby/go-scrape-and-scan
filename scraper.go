package main

import (
	"flag"
	"fmt"
	apiCalls "go-scrape-and-scan/utils/apiCalls"
	helper "go-scrape-and-scan/utils/helper"
	"os"

	"github.com/VirusTotal/vt-go"
	"github.com/gocolly/colly"
)

// flags
var url = flag.String("url", "", "url to scrape")
var apikey = flag.String("apikey", "", "VirusTotal API key")

// scraper will only visit the first page
// and will not follow any links eg. to the next page
func main() {

	flag.Parse()
	// required flag url
	if *url == "" || *apikey == "" {
		fmt.Println("url and apikey are required")
		os.Exit(0)
	}

	// Instantiate default collector
	c := colly.NewCollector()

	var links []string

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		links = append(links, h.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping
	c.Visit(*url)
	links = helper.Filter(links, *url)
	fmt.Println("Links found:", len(links))
	fmt.Println("Links:", links)

	// check if links are found
	if len(links) <= 0 {
		fmt.Println("No links found")
		os.Exit(0)
	}

	vt_client := vt.NewClient(*apikey)
	scanner := vt_client.NewURLScanner()
	fmt.Println(scanner)

	// check daily quota
	allowed, used := apiCalls.GetDailyQuota(*apikey)
	remaining := allowed - used
	fmt.Println("Daily quota for apikey:", allowed, "allowed,", used,
		"used,", remaining, "remaining")

	// check that api is capable of scanning the url list
	if remaining == 0 {
		fmt.Println("Daily quota exceeded")
		os.Exit(0)
	} else if remaining < len(links) {
		fmt.Println("Not enough quota to scan all links")
		os.Exit(0)
	}

	//report, err := scanner.Scan("https://github.com/VirusTotal/vt-go/issues/24")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(report)
}
