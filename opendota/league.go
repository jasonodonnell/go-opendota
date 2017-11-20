package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

// LeagueService provides methods for accesing information
// about leagues.
type LeagueService struct {
	sling *sling.Sling
}

func newLeagueService(sling *sling.Sling) *LeagueService {
	return &LeagueService{
		sling: sling.Path("leagues"),
	}
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
func (s *LeagueService) Leagues() ([]League, *http.Response, error) {
	leagues := new([]League)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(leagues, apiError)
	return *leagues, resp, relevantError(err, *apiError)
}
