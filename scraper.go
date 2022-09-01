package main

import (
	"flag"
	"fmt"
	apiCalls "go-scrape-and-scan/utils/apiCalls"
	helper "go-scrape-and-scan/utils/helper"
	"os"

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

	// check if links are found
	// then print amount to cmd
	if len(links) <= 0 {
		fmt.Println("No links found")
		os.Exit(0)
	}
	fmt.Println("Links found:", len(links))

	//vt_client := vt.NewClient(*apikey)
	//scanner := vt_client.NewURLScanner()
	//fmt.Println(scanner) // temp print

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
		fmt.Println("Links found exceeds daily quota")
		os.Exit(0)
	}

	// u-0908bb191bccd4d75bc210482c32633530d0e7dd3a3af3c8f1947b1a43eac817-1662056602
	//report, err := scanner.Scan("https://github.com/VirusTotal/vt-go/issues/24")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(report)
}
