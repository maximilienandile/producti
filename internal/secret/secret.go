package secret

type Algolia struct {
	AppID     string `json:"appId"`
	ApiKey    string `json:"apiKey"`
	IndexName string `json:"indexName"`
}

type Parameters struct {
	Algolia Algolia `json:"algolia"`
}
