package responseStruct

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
