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

// LeagueService provides methods for accesing information
// about leagues.
type LeagueService struct {
	sling *sling.Sling
}

// League is a collection of information about a league.
type League struct {
	LeagueID int    `json:"leagueid"`
	Ticket   string `json:"ticket"`
	Banner   string `json:"banner"`
	Tier     string `json:"tier"`
	Name     string `json:"name"`
}

// Leagues returns a collection of information about all leagues.
// https://docs.opendota.com/#tag/leagues%2Fpaths%2F~1leagues%2Fget
func (s *LeagueService) Leagues() ([]League, *http.Response, error) {
	leagues := new([]League)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(leagues, apiError)
	return *leagues, resp, relevantError(err, *apiError)
}
