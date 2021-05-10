package indexing

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"

	"github.com/maximilienandile/producti/internal/product"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	"github.com/maximilienandile/producti/internal/mocks"
)

func TestAlgoliaProductIndexer_AddProduct(t *testing.T) {
	// setup mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockAlgoliaIndexer(ctrl)
	// define product to index
	toIndex := product.Indexed{
		ProductID: "455584",
		Name:      "Leather bag",
	}
	// set mock expectations
	m.EXPECT().SaveObject(&toIndex).Return(search.SaveObjectRes{}, nil)
	// build indexer
	indexer := AlgoliaProductIndexer{indexClient: m}
	err := indexer.AddProduct(&toIndex)
	assert.Nil(t, err)
}

func TestAlgoliaProductIndexer_SearchProductByName(t *testing.T) {
	// setup mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockAlgoliaIndexer(ctrl)
	nameSearched := "Tot Bag"
	fakeRes := search.QueryRes{
		Hits: []map[string]interface{}{
			{
				"objectID": "1245",
				"name":     "Leather Bag",
			},
			{
				"objectID": "12456",
				"name":     "Tot Bag",
			},
		},
		HitsPerPage: 1000,
		NbHits:      2,
		NbPages:     1,
	}
	m.EXPECT().Search(nameSearched, opt.HitsPerPage(1000)).Return(fakeRes, nil)
	// build indexer
	indexer := AlgoliaProductIndexer{indexClient: m}
	actual, err := indexer.SearchProductByName(nameSearched)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(actual))
	assert.Contains(t, actual, &product.Indexed{
		ProductID: "1245",
		Name:      "Leather Bag",
	})
	assert.Contains(t, actual, &product.Indexed{
		ProductID: "12456",
		Name:      "Tot Bag",
	})
}
