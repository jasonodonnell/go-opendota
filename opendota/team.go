package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

type TeamService struct {
	sling *sling.Sling
}

//type TopGamesParams struct {
//	Limit  int64 `url:"limit,omitempty"`
//	Offset int64 `url:"offset,omitempty"`
//}

func newTeamService(sling *sling.Sling) *TeamService {
	return &TeamService{
		sling: sling.Path("teams"),
	}
}

type Team struct {
	TeamID        int     `json:"team_id"`
	Rating        float64 `json:"rating"`
	Wins          int     `json:"wins"`
	Losses        int     `json:"losses"`
	LastMatchTime int     `json:"last_match_time"`
	Name          string  `json:"name"`
	Tag           string  `json:"tag"`
	LogoURL       string  `json:"logo_url"`
}

func (s *TeamService) Teams() ([]Team, *http.Response, error) {
	teams := new([]Team)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(teams, apiError)
	return *teams, resp, relevantError(err, *apiError)
}
