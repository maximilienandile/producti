package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maximilienandile/producti/internal/secret"

	"github.com/golang/mock/gomock"
	"github.com/maximilienandile/producti/internal/brand"
	"github.com/maximilienandile/producti/internal/mocks"
	"github.com/maximilienandile/producti/internal/price"
	"github.com/maximilienandile/producti/internal/product"
)

var testProduct product.Product
var testProductJSON string

func init() {
	testProduct = product.Product{
		Name:                       "test",
		OriginalPrice:              price.Price{},
		Brand:                      brand.Brand{},
		Followers:                  0,
		DaysOnline:                 0,
		ViewsSinceLastWeek:         0,
		IsPriceDropAlertView:       false,
		IsPriceDropAlertDaysOnline: false,
		Pictures:                   nil,
		PriceDropped:               price.Price{},
		RecommendedPrice:           price.Price{},
	}
	productMarshalled, err := json.Marshal(testProduct)
	if err != nil {
		panic(err)
	}
	testProductJSON = string(productMarshalled)
}

func TestCreateProductOK(t *testing.T) {
	// setup mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockProductStore(ctrl)
	m.EXPECT().Create(&testProduct).Return(&testProduct, nil)

	// build test server
	conf := Config{
		Secrets:      secret.Parameters{},
		ProductStore: m,
	}
	testServer := New(&conf)
	// create response recorder
	w := httptest.NewRecorder()
	// build test request
	var buf bytes.Buffer
	buf.WriteString(testProductJSON)
	testRequest, err := http.NewRequest("POST", "/product", &buf)
	assert.Nil(t, err)
	// send the test request
	testServer.GinEngine.ServeHTTP(w, testRequest)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, testProductJSON, w.Body.String())
}

func TestCreateProductBadRequest(t *testing.T) {
	// setup mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockProductStore(ctrl)
	// build test server
	conf := Config{
		Secrets:      secret.Parameters{},
		ProductStore: m,
	}
	testServer := New(&conf)
	// create response recorder
	w := httptest.NewRecorder()
	// build test request
	var buf bytes.Buffer
	buf.WriteString(`yo`)
	testRequest, err := http.NewRequest("POST", "/product", &buf)
	assert.Nil(t, err)
	// send the test request
	testServer.GinEngine.ServeHTTP(w, testRequest)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
