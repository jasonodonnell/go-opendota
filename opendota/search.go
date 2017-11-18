package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

// SearchService provides methods for searching players by
// personaname.  Default similarity is 0.51.
type SearchService struct {
	sling *sling.Sling
}

// Search is a collection about a player.
type Search struct {
	AccountID     int     `json:"account_id"`
	AvatarFull    string  `json:"avatarfull"`
	Personaname   string  `json:"personaname"`
	LastMatchTime string  `json:"last_match_time"`
	Similarity    float64 `json:"similarity"`
}

// SearchParams are the paramters for querying the
// search service.
type SearchParams struct {
	Query      string  `url:"q"`
	Similarity float64 `url:"similarity,omitempty"`
}

func newSearchService(sling *sling.Sling) *SearchService {
	return &SearchService{
		sling: sling.Path("search"),
	}
}

// Search returns an array of players who are similar to the query
// provided.
func (s *SearchService) Search(params *SearchParams) ([]Search, *http.Response, error) {
	search := new([]Search)
	apiError := new(APIError)
	resp, err := s.sling.New().QueryStruct(params).Receive(search, apiError)
	return *search, resp, relevantError(err, *apiError)
}
