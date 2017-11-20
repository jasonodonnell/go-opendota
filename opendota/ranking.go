package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

// RankingService provides methods for accessing the ranking of
// heroes by players.
type RankingService struct {
	sling *sling.Sling
}

// Ranking is a collection of information about the top
// players of a specific hero.
type Ranking struct {
	HeroID   int       `json:"hero_id"`
	Rankings []ranking `json:"rankings"`
}

type ranking struct {
	AccountID           int     `json:"account_id"`
	Score               float64 `json:"score"`
	Personaname         string  `json:"personaname"`
	Name                string  `json:"name"`
	Avatar              string  `json:"avatar"`
	LastLogin           string  `json:"last_login"`
	SoloCompetitiveRank int     `json:"solo_competitive_rank"`
}

// RankingParam is the parameter to search the rankings of
// a specific hero.
type RankingParam struct {
	HeroID string `url:"hero_id"`
}

func newRankingService(sling *sling.Sling) *RankingService {
	return &RankingService{
		sling: sling.Path("rankings"),
	}
}

// Rankings returns the top ranking of a hero by players.
func (s *RankingService) Rankings(param *RankingParam) (Ranking, *http.Response, error) {
	rankings := new(Ranking)
	apiError := new(APIError)
	resp, err := s.sling.New().QueryStruct(param).Receive(rankings, apiError)
	return *rankings, resp, relevantError(err, *apiError)
}
