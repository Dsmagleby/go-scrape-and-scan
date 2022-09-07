package dataStruct

type Quota struct {
	Data struct {
		APIRequestsDaily struct {
			Group struct {
				Allowed       int    `json:"allowed"`
				InheritedFrom string `json:"inherited_from"`
				Used          int    `json:"used"`
			} `json:"group"`
			User struct {
				Allowed int `json:"allowed"`
				Used    int `json:"used"`
			} `json:"user"`
		} `json:"api_requests_daily"`
		APIRequestsHourly struct {
			Group struct {
				Allowed       int    `json:"allowed"`
				InheritedFrom string `json:"inherited_from"`
				Used          int    `json:"used"`
			} `json:"group"`
			User struct {
				Allowed int64 `json:"allowed"`
				Used    int   `json:"used"`
			} `json:"user"`
		} `json:"api_requests_hourly"`
		APIRequestsMonthly struct {
			Group struct {
				Allowed       int    `json:"allowed"`
				InheritedFrom string `json:"inherited_from"`
				Used          int    `json:"used"`
			} `json:"group"`
			User struct {
				Allowed int `json:"allowed"`
				Used    int `json:"used"`
			} `json:"user"`
		} `json:"api_requests_monthly"`
	} `json:"data"`
}

type AnalysisReport struct {
	Data struct {
		Attributes struct {
			LastHTTPResponseContentSha256 string        `json:"last_http_response_content_sha256"`
			LastFinalURL                  string        `json:"last_final_url"`
			LastHTTPResponseContentLength int           `json:"last_http_response_content_length"`
			URL                           string        `json:"url"`
			LastAnalysisDate              int           `json:"last_analysis_date"`
			Tags                          []interface{} `json:"tags"`
			LastAnalysisStats             struct {
				Harmless   int `json:"harmless"`
				Malicious  int `json:"malicious"`
				Suspicious int `json:"suspicious"`
				Undetected int `json:"undetected"`
				Timeout    int `json:"timeout"`
			} `json:"last_analysis_stats"`
			Reputation           int      `json:"reputation"`
			LastModificationDate int      `json:"last_modification_date"`
			Title                string   `json:"title"`
			OutgoingLinks        []string `json:"outgoing_links"`
			TimesSubmitted       int      `json:"times_submitted"`
			FirstSubmissionDate  int      `json:"first_submission_date"`
			TotalVotes           struct {
				Harmless  int `json:"harmless"`
				Malicious int `json:"malicious"`
			} `json:"total_votes"`
		} `json:"attributes"`
		Type  string `json:"type"`
		ID    string `json:"id"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
}
