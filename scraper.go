package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"utils/responseStruct"

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
		links = append(links, link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping
	c.Visit(*url)

	vt_client := vt.NewClient(*apikey)
	scanner := vt_client.NewURLScanner()

	// check daily quota
	client := http.Client{}
	quotaEndpoint := "https://www.virustotal.com/api/v3/users/" + *apikey + "/overall_quotas"
	req, err := http.NewRequest("GET", quotaEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Apikey", *apikey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}

	var quota responseStruct.Quota
	if err := json.NewDecoder(resp.Body).Decode(&quota); err != nil {
		log.Fatal(err)
	}
	fmt.Println(quota.Data.APIRequestsDaily.User.Used)

	report, err := scanner.Scan("https://github.com/VirusTotal/vt-go/issues/24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(report)
}
