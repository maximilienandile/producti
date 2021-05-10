package indexing

import (
	"errors"
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/maximilienandile/producti/internal/product"
	"github.com/maximilienandile/producti/internal/secret"
)

const maxHitsPerPageAlgolia = 1000

type AlgoliaProductIndexer struct {
	indexClient AlgoliaIndexer
}

func NewAlgoliaProductIndexer(algoliaConf secret.Algolia) ProductIndexer {
	return AlgoliaProductIndexer{
		indexClient: search.NewClient(algoliaConf.AppID, algoliaConf.APIKey).InitIndex(algoliaConf.IndexName),
	}
}

func (a AlgoliaProductIndexer) AddProduct(product *product.Indexed) error {
	// will save an object into the index making it available for search
	_, err := a.indexClient.SaveObject(product)
	return err
}

func (a AlgoliaProductIndexer) SearchProductByName(name string) ([]*product.Indexed, error) {
	// TODO : when other fields get indexed limit the search to the attribute name only
	// we ask for the maximum number of hists to avoid new paginated requests
	res, err := a.indexClient.Search(name, opt.HitsPerPage(maxHitsPerPageAlgolia))
	if err != nil {
		return nil, err
	}
	productsFound := make([]*product.Indexed, 0, res.NbHits)
	for _, hit := range res.Hits {

		objectID, found := hit["objectID"]
		if !found {
			return nil, errors.New("cannot unmarshall results of search no objectID")
		}
		name, found := hit["name"]
		if !found {
			return nil, errors.New("cannot unmarshall results of search no name")
		}
		productsFound = append(productsFound, &product.Indexed{
			ProductID: fmt.Sprintf("%v", objectID),
			Name:      fmt.Sprintf("%v", name),
		})
	}
	return productsFound, nil
}
