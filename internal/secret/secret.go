package secret

type Algolia struct {
	AppID     string `json:"appId"`
	APIKey    string `json:"apiKey"`
	IndexName string `json:"indexName"`
}

type Parameters struct {
	Algolia Algolia `json:"algolia"`
}
