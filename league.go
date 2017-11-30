package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

func newLeagueService(sling *sling.Sling) *LeagueService {
	return &LeagueService{
		sling: sling.Path("leagues"),
	}
}

// LeagueService provides a method for accessing information
// about leagues.
type LeagueService struct {
	sling *sling.Sling
}

// League represents a league in Dota 2.
type League struct {
	LeagueID int    `json:"leagueid"`
	Ticket   string `json:"ticket"`
	Banner   string `json:"banner"`
	Tier     string `json:"tier"`
	Name     string `json:"name"`
}

// Leagues returns a collection of all leagues in Dota 2.
// https://docs.opendota.com/#tag/leagues%2Fpaths%2F~1leagues%2Fget
func (s *LeagueService) Leagues() ([]League, *http.Response, error) {
	leagues := new([]League)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(leagues, apiError)
	return *leagues, resp, relevantError(err, *apiError)
}
