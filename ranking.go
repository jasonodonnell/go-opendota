package opendota

import (
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

func newRankingService(sling *sling.Sling) *RankingService {
	return &RankingService{
		sling: sling.Path("rankings"),
	}
}

// RankingService provides a method for accessing ranking of
// heroes for a player.
type RankingService struct {
	sling *sling.Sling
}

type rankingParam struct {
	heroID string `url:"hero_id"`
}

// Ranking represents the top player rankings for a hero.
type HeroRanking struct {
	HeroID   int       `json:"hero_id"`
	Rankings []Ranking `json:"rankings"`
}

type Ranking struct {
	AccountID           int     `json:"account_id"`
	Score               float64 `json:"score"`
	Personaname         string  `json:"personaname"`
	Name                string  `json:"name"`
	Avatar              string  `json:"avatar"`
	LastLogin           string  `json:"last_login"`
	SoloCompetitiveRank int     `json:"solo_competitive_rank"`
}

// Rankings takes a Hero ID and returns the top player rankings for a hero.
// https://docs.opendota.com/#tag/rankings%2Fpaths%2F~1rankings%2Fget
func (s *RankingService) Rankings(heroID int) (HeroRanking, *http.Response, error) {
	params := &rankingParam{}
	params.heroID = strconv.Itoa(heroID)
	rankings := new(HeroRanking)
	apiError := new(APIError)
	resp, err := s.sling.New().QueryStruct(params).Receive(rankings, apiError)
	return *rankings, resp, relevantError(err, *apiError)
}
