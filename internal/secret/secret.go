package secret

// Secrets needed to communicate with algolia.
type Algolia struct {
	AppID     string `json:"appId"`
	APIKey    string `json:"apiKey"`
	IndexName string `json:"indexName"`
}

// Secret parameters.
// Add here any secret.
type Parameters struct {
	Algolia Algolia `json:"algolia"`
}
