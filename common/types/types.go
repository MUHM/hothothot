package baseTypes

type PasswordConfig struct {
	Secret  string `json:"secret"`
	Default string `json:"default"`
}

type CrawlerZhihu struct {
	URL       string `json:"url"`
	Accept    string `json:"accept"`
	UserAgent string `json:"userAgent"`
}
