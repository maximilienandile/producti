package price

type Price struct {
	Currency  string `json:"currency"`
	Cents     int    `json:"cents"`
	Formatted string `json:"formatted"`
}
