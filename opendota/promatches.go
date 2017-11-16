package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

func newProMatchesService(sling *sling.Sling) *ProMatchesService {
	return &ProMatchesService{
		sling: sling.Path("proMatches"),
	}
}

// ProMatchesService provides methods for accessing pro
// matches.
type ProMatchesService struct {
	sling *sling.Sling
}

// ProMatch is a collection about a pro match.
type ProMatch struct {
	MatchID       int64  `json:"match_id"`
	Duration      int    `json:"duration"`
	StartTime     int    `json:"start_time"`
	RadiantTeamID int    `json:"radiant_team_id"`
	RadiantName   string `json:"radiant_name"`
	DireTeamID    int    `json:"dire_team_id"`
	DireName      string `json:"dire_name"`
	LeagueID      int    `json:"leagueid"`
	LeagueName    string `json:"league_name"`
	SeriesID      int    `json:"series_id"`
	SeriesType    int    `json:"series_type"`
	RadiantScore  int    `json:"radiant_score"`
	DireScore     int    `json:"dire_score"`
	RadiantWin    bool   `json:"radiant_win"`
}

// Matches returns information about pro matches.
func (s *ProMatchesService) Matches() ([]ProMatch, *http.Response, error) {
	promatches := new([]ProMatch)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(promatches, apiError)
	return *promatches, resp, relevantError(err, *apiError)
}
