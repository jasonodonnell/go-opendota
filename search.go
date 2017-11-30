package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

func newSearchService(sling *sling.Sling) *SearchService {
	return &SearchService{
		sling: sling.Path("search"),
	}
}

// SearchService provides a method for searching players by
// personaname.
type SearchService struct {
	sling *sling.Sling
}

// SearchParams represents optional query parameters for the Search method.
// Default similarity is 0.51.
type SearchParams struct {
	query      string  `url:"q"`
	Similarity float64 `url:"similarity,omitempty"`
}

// Search represents a player for a given personaname.
type Search struct {
	AccountID     int     `json:"account_id"`
	AvatarFull    string  `json:"avatarfull"`
	Personaname   string  `json:"personaname"`
	LastMatchTime string  `json:"last_match_time"`
	Similarity    float64 `json:"similarity"`
}

// Search takes a query string and optional params and returns an array
// of players who are similar to the query provided.
// https://docs.opendota.com/#tag/search%2Fpaths%2F~1search%2Fget
func (s *SearchService) Search(query string, params *SearchParams) ([]Search, *http.Response, error) {
	if params == nil {
		params = &SearchParams{}
	}
	params.query = query
	search := new([]Search)
	apiError := new(APIError)
	resp, err := s.sling.New().QueryStruct(params).Receive(search, apiError)
	return *search, resp, relevantError(err, *apiError)
}
