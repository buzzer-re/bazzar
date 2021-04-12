package abuse

import "net/url"

const (
	URLHAUS_API_URL     = "https://urlhaus-api.abuse.ch/v1/url"
	URLHAUS_API_HOST    = "https://urlhaus-api.abuse.ch/v1/host/"
	URLHAUS_API_PAYLOAD = "https://urlhaus-api.abuse.ch/v1/payload/"
	URLHAUS_API_TAG     = "https://urlhaus-api.abuse.ch/v1/tag/"
)

var (
	query_url_post  url.Values = url.Values{"url": {}}
	query_host_post url.Values = url.Values{"host": {}}
)

type URLResponse struct {
	QueryStatus      string `json:"query_status"`
	ID               string `json:"id"`
	UrlhausReference string `json:"urlhaus_reference"`
	URL              string `json:"url"`
	URLStatus        string `json:"url_status"`
	Host             string `json:"host"`
	DateAdded        string `json:"date_added"`
	Threat           string `json:"threat"`
	Blacklists       struct {
		SpamhausDbl string `json:"spamhaus_dbl"`
		Surbl       string `json:"surbl"`
	} `json:"blacklists"`
	Reporter            string   `json:"reporter"`
	Larted              string   `json:"larted"`
	TakedownTimeSeconds string   `json:"takedown_time_seconds"`
	Tags                []string `json:"tags"`
	Payloads            []struct {
		Firstseen       string `json:"firstseen"`
		Filename        string `json:"filename"`
		FileType        string `json:"file_type"`
		ResponseSize    string `json:"response_size"`
		ResponseMd5     string `json:"response_md5"`
		ResponseSha256  string `json:"response_sha256"`
		UrlhausDownload string `json:"urlhaus_download"`
		Signature       string `json:"signature"`
		Virustotal      struct {
			Result  string `json:"result"`
			Percent string `json:"percent"`
			Link    string `json:"link"`
		} `json:"virustotal"`
		Imphash string `json:"imphash"`
		Ssdeep  string `json:"ssdeep"`
		Tlsh    string `json:"tlsh"`
	} `json:"payloads"`
}

type HostResponse struct {
	QueryStaus       string `json:"query_staus"`
	UrlhausReference string `json:"urlhaus_reference"`
	Host             string `json:"host"`
	Firstseen        string `json:"firstseen"`
	URLCount         string `json:"url_count"`
	Blacklists       struct {
		SpamhausDbl string `json:"spamhaus_dbl"`
		Surbl       string `json:"surbl"`
	} `json:"blacklists"`
	Urls []struct {
		ID                  string      `json:"id"`
		UrlhausReference    string      `json:"urlhaus_reference"`
		URL                 string      `json:"url"`
		URLStatus           string      `json:"url_status"`
		DateAdded           string      `json:"date_added"`
		Threat              string      `json:"threat"`
		Reporter            string      `json:"reporter"`
		Larted              string      `json:"larted"`
		TakedownTimeSeconds interface{} `json:"takedown_time_seconds"`
		Tags                []string    `json:"tags"`
	} `json:"urls"`
}
