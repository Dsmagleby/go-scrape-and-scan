package apiCalls

import (
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
