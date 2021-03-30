package abuse

import "net/url"

// Constants from Malware bazzar
const (
	MALWARE_BAZZAR_API_URL = "https://mb-api.abuse.ch/api/v1/"
	MB_RESPONSE_OK         = "ok"
	UNKNOWN_FILETYPE       = "unknown"
)

// Post form variables
var (
	getSampleInfoForm     url.Values = url.Values{"query": {"get_info"}, "hash": {}}
	sampleByTag           url.Values = url.Values{"query": {"get_taginfo"}, "tag": {}, "limit": {"1000"}}
	getSample			  url.Values = url.Values{"query": {"get_file"}, "sha256_hash": {}}
	latestSamplesFormData url.Values = url.Values{"query": {"get_recent"}, "selector": {"time"}}
	querySampleSignature  url.Values = url.Values{"query": {"get_siginfo"}, "signature": {}, "limit": {}}
	// queryClamavSignature  url.Values = url.Values{"query": {"get_clamavinfo"}, "clamav": {}, "limit": {}}
)

// Struct of each api response

// Sample data generic struct
type Data struct {
	Sha256Hash    string      `json:"sha256_hash"`
	Sha3384Hash   string      `json:"sha3_384_hash"`
	Sha1Hash      string      `json:"sha1_hash"`
	Md5Hash       string      `json:"md5_hash"`
	FirstSeen     string      `json:"first_seen"`
	LastSeen      string      `json:"last_seen"`
	FileName      string      `json:"file_name"`
	FileSize      int         `json:"file_size"`
	FileTypeMime  string      `json:"file_type_mime"`
	FileType      string      `json:"file_type"`
	Reporter      string      `json:"reporter"`
	OriginCountry string      `json:"origin_country"`
	Anonymous     int         `json:"anonymous"`
	Signature     string      `json:"signature"`
	Imphash       string      `json:"imphash"`
	Tlsh          string      `json:"tlsh"`
	Ssdeep        string      `json:"ssdeep"`
	Tags          []string    `json:"tags"`
	CodeSign      string `json:"code_sign"`
	Intelligence  struct {
		Clamav    []string `json:"clamav"`
		Downloads string   `json:"downloads"`
		Uploads   string   `json:"uploads"`
		Mail      struct {
			Generic string `json:"Generic"`
			CH      string `json:"CH"`
		} `json:"mail"`
	} `json:"intelligence"`
}

type SampleInfo struct {
	Sha256Hash     string      `json:"sha256_hash"`
	Sha3384Hash    string      `json:"sha3_384_hash"`
	Sha1Hash       string      `json:"sha1_hash"`
	Md5Hash        string      `json:"md5_hash"`
	FirstSeen      string      `json:"first_seen"`
	LastSeen       string      `json:"last_seen"`
	FileName       string      `json:"file_name"`
	FileSize       int         `json:"file_size"`
	FileTypeMime   string      `json:"file_type_mime"`
	FileType       string      `json:"file_type"`
	Reporter       string      `json:"reporter"`
	OriginCountry  string      `json:"origin_country"`
	Anonymous      int         `json:"anonymous"`
	Signature      string      `json:"signature"`
	Imphash        string      `json:"imphash"`
	Tlsh           string      `json:"tlsh"`
	Ssdeep         string      `json:"ssdeep"`
	Comment        string      `json:"comment"`
	Tags           []string    `json:"tags"`
	CodeSign       string `json:"code_sign"`
	DeliveryMethod string      `json:"delivery_method"`
	Intelligence   struct {
		Clamav    []string    `json:"clamav"`
		Downloads string      `json:"downloads"`
		Uploads   string      `json:"uploads"`
		Mail      interface{} `json:"mail"`
	} `json:"intelligence"`
	FileInformation []struct {
		Context string `json:"context"`
		Value   string `json:"value"`
	} `json:"file_information"`
	// OleInformation struct{ THIS FIELD GIVE US TOO MUCH NOISE AND CHANGE OVER SAMPLE
	// 	Olevba []struct {
	// 		Type string `json:"type"`
	// 		Keyword string `json:"Keyword"`
	// 		Description string `json:"description"`
	// 	}
	// } `json:"ole_information,omitempty"`
	YaraRules      []struct{
		RuleName	string `json:"rule_name"`
		Author		string `json:"author"`
		Description	string `json:"description"`
		Reference	string `json:"reference"`
	} `json:"yara_rules"`
	VendorIntel    struct {
		CERTPLMWDB struct {
			Detection string `json:"detection"`
			Link      string `json:"link"`
		} `json:"CERT-PL_MWDB"`
		YOROIYOMI struct {
			Detection string `json:"detection"`
			Score     string `json:"score"`
		} `json:"YOROI_YOMI"`
		VxCube struct {
			Verdict       string `json:"verdict"`
			Maliciousness string `json:"maliciousness"`
			Behaviour     []struct {
				ThreatLevel string `json:"threat_level"`
				Rule        string `json:"rule"`
			} `json:"behaviour"`
		} `json:"vxCube"`
		CAPE struct {
			Detection string `json:"detection"`
			Link      string `json:"link"`
		} `json:"CAPE"`
		Triage struct {
			MalwareFamily string   `json:"malware_family"`
			Score         string   `json:"score"`
			Link          string   `json:"link"`
			Tags          []string `json:"tags"`
			Signatures    []struct {
				Signature string `json:"signature"`
				Score     string `json:"score"`
			} `json:"signatures"`
			MalwareConfig []struct {
				Extraction string `json:"extraction"`
				Family     string `json:"family"`
				C2         string `json:"c2"`
			} `json:"malware_config"`
		} `json:"Triage"`
		ReversingLabs struct {
			ThreatName     string `json:"threat_name"`
			Status         string `json:"status"`
			FirstSeen      string `json:"first_seen"`
			ScannerCount   string `json:"scanner_count"`
			ScannerMatch   string `json:"scanner_match"`
			ScannerPercent string `json:"scanner_percent"`
		} `json:"ReversingLabs"`
		SpamhausHBL []struct {
			Detection string `json:"detection"`
			Link      string `json:"link"`
		} `json:"Spamhaus_HBL"`
		UnpacMe []struct {
			Sha256Hash string        `json:"sha256_hash"`
			Md5Hash    string        `json:"md5_hash"`
			Sha1Hash   string        `json:"sha1_hash"`
			Detections []string `json:"detections"`
			Link       string        `json:"link"`
		} `json:"UnpacMe"`
	} `json:"vendor_intel"`
}

type SampleQuery struct {
	QueryStatus string       `json:"query_status"`
	Data        []SampleInfo `json:"data"`
}
type Response struct {
	QueryStatus string `json:"query_status"`
	Data        []Data
}

type QueryStatus struct {
	QueryStatus string `json:"query_status"`
}
