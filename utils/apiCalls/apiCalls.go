package apiCalls

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	dataStruct "go-scrape-and-scan/utils/dataStruct"
)

// call endpoint to get available daily quota
func GetDailyQuota(apikey string) (int, int) {
	client := http.Client{}
	quotaEndpoint := "https://www.virustotal.com/api/v3/users/" + apikey + "/overall_quotas"
	req, err := http.NewRequest("GET", quotaEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Apikey", apikey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// extract daily quota from response
	var quota dataStruct.Quota
	json.NewDecoder(resp.Body).Decode(&quota)
	return quota.Data.APIRequestsDaily.User.Allowed, quota.Data.APIRequestsDaily.User.Used
}

func ScanUrl(apikey string, url string) {
	client := http.Client{}
	scanEndpoint := "https://www.virustotal.com/api/v3/urls"
	req, err := http.NewRequest("POST", scanEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Apikey", apikey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func GetReport(apikey string, url string) []int {
	client := http.Client{}
	// generate URL identifier
	var urlID = base64.RawURLEncoding.EncodeToString([]byte(url))
	reportEndpoint := "https://www.virustotal.com/api/v3/urls/" + urlID
	req, err := http.NewRequest("GET", reportEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Apikey", apikey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var report dataStruct.AnalysisReport
	json.NewDecoder(resp.Body).Decode(&report)

	var results []int
	results = append(results, report.Data.Attributes.LastAnalysisStats.Harmless)
	results = append(results, report.Data.Attributes.LastAnalysisStats.Malicious)
	results = append(results, report.Data.Attributes.LastAnalysisStats.Suspicious)
	results = append(results, report.Data.Attributes.LastAnalysisStats.Undetected)
	return results
}
