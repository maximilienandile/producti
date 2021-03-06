package product

import (
	"github.com/maximilienandile/producti/internal/brand"
	"github.com/maximilienandile/producti/internal/picture"
	"github.com/maximilienandile/producti/internal/price"
)

type Product struct {
	ID                         string            `json:"id"`
	Name                       string            `json:"name"`
	OriginalPrice              price.Price       `json:"originalPrice"`
	Brand                      brand.Brand       `json:"brand"`
	Followers                  uint              `json:"followers"`
	DaysOnline                 uint              `json:"daysOnline"`
	ViewsSinceLastWeek         uint              `json:"viewsSinceLastWeek"`
	IsPriceDropAlertView       bool              `json:"isPriceDropAlertView"`
	IsPriceDropAlertDaysOnline bool              `json:"isPriceDropAlertDaysOnline"`
	Pictures                   []picture.Picture `json:"pictures"`
	PriceDropped               price.Price       `json:"priceDropped"`
	RecommendedPrice           price.Price       `json:"recommendedPrice"`
}

// Indexed is the product representation indexed.
// Note that just the fields indexed are in that struct.
// It is made to make the index lighter (we do not want fields in the index that are not searchable).
type Indexed struct {
	ProductID string `json:"objectID"`
	Name      string `json:"name"`
}
