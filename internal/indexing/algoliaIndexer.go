package indexing

import "github.com/algolia/algoliasearch-client-go/v3/algolia/search"

// interface that lists methods used to interact with algolia
// this interface is used to facilitate unit testing
type AlgoliaIndexer interface {
	SaveObject(object interface{}, opts ...interface{}) (res search.SaveObjectRes, err error)
	Search(query string, opts ...interface{}) (res search.QueryRes, err error)
}
